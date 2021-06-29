package dB

import (
	"context"
	"fmt"
	"log"
	pb "mobileapps/jobsserver/protos/gen/jobslist"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

const firebase_config_file = "../config/bangalore-jobs-app-firebase-adminsdk-z4dti-3bbb334a23.json"

func GetJobs() []*pb.Job {
	var job_response []*pb.Job
	ctx := context.Background()
	config := &firebase.Config{DatabaseURL: "https://bangalore-jobs-app.firebaseio.com"}

	opt := option.WithCredentialsFile(firebase_config_file)
	app, err := firebase.NewApp(ctx, config, opt)

	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		//return job_response
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ref := client.NewRef("/jobs").OrderByKey().LimitToLast(100)
	var data map[string]map[string]pb.Job

	if err := ref.Get(ctx, &data); err != nil {
		log.Fatalln("Error reading from database:", err)
	}
	for _, val := range data {

		for _, job_obj := range val {
			job_response = append(job_response, &job_obj)
		}

	}

	return job_response
}
