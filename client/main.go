package main 

import(
	"context"
	"grpc"
	"log"
	pb "/20summer/project_one/projectone.pb.go"
)

func main(){
	conn,err:=grpc.Dial("localhost:50051")
	if err!=nil {
		log.Fatalf("did not connect:%v",err)
	}
	defer conn.close()
	c:=pb.NewDeal_CardClient()
	ctx,cancel:=context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err:=c.DealCard(ctx,&pb.Name{Name:"mike"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s:%s", r.GetName(),r.GetCard())

}