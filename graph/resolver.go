package graph

import "go-graphql/graph/model"

type Resolver struct {
	CharacterStore map[string]model.Character
}
