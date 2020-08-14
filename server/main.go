package main

import (
	"fmt"
	"./Deck"
	"math/rand"
	"time"
	"context"
	"net"
	"log"
	"grpc"
	pb "~/20summer/project_one/projectone.pb.go"
)

type Server struct{
	pb.UnimplementedDeal_CardServer
}

type Person struct{
	string name
	string card
}

type Name struct{
	string name
}

fun (S *Server) DealCard (ctx context.context, name *Name) (*Person,err){
	rand.Seed(time.Now().Unix())
	card_one:=rand.Intn(52)+1
	person:=new(Person)
	person.name=name.name
	person.card=Deck.Card_map[card_one]
	return person,nil
}

func main(){
	lis,err:=net.Listen("tcp",":50051")
	if err!=nil {
		log.Fatalf("fail to listen: %v",err)
	}
	s:=grpc.NewServer()
	pb.RegisterDeal_CardServer(s,&Server{})
	grpcServer.Serve(lis)
}