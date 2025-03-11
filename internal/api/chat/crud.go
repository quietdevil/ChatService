package chat

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/quietdevil/ChatSevice/internal/logger"
	desc "github.com/quietdevil/ChatSevice/pkg/chat_v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ImplementationChat) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	logger.Info("CreateChat", zap.Any("idU", id.ID()))

	idStr := strconv.FormatUint(uint64(id.ID()), 10)
	s.channels[idStr] = make(chan *desc.Message, 100)
	//id, err := s.ChatService.Create(ctx, req.Usernames)
	//if err != nil {
	//	return nil, err
	//}
	return &desc.CreateResponse{
		Id: int64(id.ID()),
	}, nil
}

func (s *ImplementationChat) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.ChatService.Delete(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *ImplementationChat) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	s.mxChannels.RLock()
	chatChan, ok := s.channels[req.GetId()]
	s.mxChannels.RUnlock()
	if !ok {
		return nil, errors.New("chan not found")
	}

	chatChan <- &desc.Message{
		From:      req.Message.GetFrom(),
		Text:      req.Message.GetText(),
		Timestamp: timestamppb.New(time.Now().Local()),
	}
	return &emptypb.Empty{}, nil
}

func (s *ImplementationChat) Connect(req *desc.ConnectRequest, stream grpc.ServerStreamingServer[desc.Message]) error {
	s.mxChannels.RLock()
	chatChan, ok := s.channels[req.GetId()]
	s.mxChannels.RUnlock()
	if !ok {
		return errors.New("There is no chat for this ID")
	}

	s.mx.Lock()
	if _, okChat := s.chats[req.GetId()]; !okChat {
		s.chats[req.GetId()] = &Chat{
			streams: make(map[string]desc.Chat_ConnectServer),
		}
	}
	s.mx.Unlock()

	s.chats[req.GetId()].m.Lock()
	s.chats[req.GetId()].streams[req.GetUsername()] = stream
	s.chats[req.GetId()].m.Unlock()
	for {
		select {
		case <-stream.Context().Done():
			s.chats[req.GetId()].m.Lock()
			delete(s.chats[req.GetId()].streams, req.GetId())
			s.chats[req.GetId()].m.Unlock()
			return errors.New("Connection closed")

		case msg, ok := <-chatChan:
			if !ok {
				return errors.New("Chat channel closed")
			}

			for _, st := range s.chats[req.GetId()].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}
		}
	}
}
