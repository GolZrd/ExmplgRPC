package my_chat

import context "context"

type MyBeautifulServer struct {
}

func (MyBeautifulServer) Hi(_ context.Context, msg *ChatMessage) (*ChatMessage, error) {
	return &ChatMessage{Text: msg.Text}, nil
}

func (MyBeautifulServer) mustEmbedUnimplementedChatServiceServer() {}
