package main

import (
	"gobblet-gobblers-go/logic"
)

type GameReponse struct {
	RoomId    string           `json:"room_id"`
	Board     BoardResponse    `json:"board"`
	P1        PlayerResponse   `json:"p1"`
	P2        PlayerResponse   `json:"p2"`
	Turn      logic.P          `json:"turn"`
	RestPiece []*PieceResponse `json:"rest_piece"`
	AllPiece  []*PieceResponse `json:"all_piece"`
	Win       *PlayerResponse  `json:"win"`
}

type PlayerResponse struct {
	Name   string  `json:"name"`
	MeTurn logic.P `json:"my_turn"`
}

type BoardResponse [logic.BOARD_MAX][logic.BOARD_MAX]PiecesResponse
type PiecesResponse [3]*PieceResponse

type PieceResponse struct {
	Id   uint       `json:"id"`
	Size logic.Size `json:"size"`
	User logic.P    `json:"user"`
}

func BoardToResponse(b logic.Board) BoardResponse {
	br := BoardResponse{}
	for x, col := range b {
		for y, val := range col {
			psr := PiecesResponse{}
			for i := 0; i < 3; i++ {
				if val.GetPieces()[i] == nil {
					break
				}
				pr := new(PieceResponse)
				pr.Id = val.GetPieces()[i].GetId()
				pr.Size = val.GetPieces()[i].GetSize()
				pr.User = val.GetPieces()[i].GetUser()
				psr[i] = pr
			}
			br[x][y] = psr
		}
	}
	return br
}

func PlayerToResponse(p *logic.Player) *PlayerResponse {
	if p == nil {
		return nil
	}
	pr := PlayerResponse{
		Name:   p.GetName(),
		MeTurn: p.GetP(),
	}
	return &pr
}

func RestPiecesToRestPiecesResponse(p []*logic.Piece) []*PieceResponse {
	ps := make([]*PieceResponse, 0)
	for _, v := range p {
		p := new(PieceResponse)
		p.Size = v.GetSize()
		p.User = v.GetUser()
		p.Id = v.GetId()
		ps = append(ps, p)
	}
	return ps
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type OkResponse struct {
	Message string `json:"message"`
}
