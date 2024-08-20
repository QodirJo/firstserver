package handler

import (
	"log"

	"github.com/Exam4/ApiGatawey/genproto/health_sync"
	"github.com/Exam4/ApiGatawey/kafka"
	"google.golang.org/grpc"
)

type Handler struct {
	Generic               health_sync.GeneticServiceClient
	Wearable              health_sync.WearableServiceClient
	Records               health_sync.MedicalRecordServiceClient
	HealthRecommendations health_sync.HealthRecommendationServiceClient
	Lifestyle             health_sync.LifestyleServiceClient
	Producer              kafka.KafkaProducer
}

func NewHandler(healthconn *grpc.ClientConn, Producer kafka.KafkaProducer) *Handler {
	pr, err := kafka.NewKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		log.Fatal(err)
	}

	return &Handler{
		Generic:               health_sync.NewGeneticServiceClient(healthconn),
		Wearable:              health_sync.NewWearableServiceClient(healthconn),
		Records:               health_sync.NewMedicalRecordServiceClient(healthconn),
		HealthRecommendations: health_sync.NewHealthRecommendationServiceClient(healthconn),
		Lifestyle:             health_sync.NewLifestyleServiceClient(healthconn),
		Producer:              pr,
	}

}
