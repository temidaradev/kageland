package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tomlister/kageland/util"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type ShaderNewResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Likes      int    `json:"likes"`
	Views      int    `json:"views"`
	FragShader string `json:"frag_shader"`
	Image1     string `json:"image_1"`
	Image2     string `json:"image_2"`
	Image3     string `json:"image_3"`
	Image4     string `json:"image_4"`
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		db = &util.DB{}
		db.Connect()
	}
	prod := db.Client.Database("prod")
	shaders := prod.Collection("shaders")
	opts := options.Find().SetSort(bson.M{"created": -1}).SetLimit(3)
	cursor, err := shaders.Find(db.Ctx, bson.M{}, opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	var topShaders []ShaderDocument
	err = cursor.All(db.Ctx, &topShaders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	resp := []ShaderTopResponse{}
	for _, s := range topShaders {
		resp = append(resp, ShaderTopResponse{
			ID:         s.ID,
			Name:       s.Name,
			Views:      s.Views,
			Likes:      s.Likes,
			FragShader: s.FragShader,
			Image1:     s.Image1,
			Image2:     s.Image2,
			Image3:     s.Image3,
			Image4:     s.Image4,
		})
	}
	json.NewEncoder(w).Encode(resp)
}
