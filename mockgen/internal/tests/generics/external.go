package generics

import (
	"context"
	"io"

	"go.uber.org/mock/mockgen/internal/tests/generics/other"
	"golang.org/x/exp/constraints"
)

//go:generate mockgen --source=external.go --destination=source/mock_external_test.go --package source

type ExternalConstraint[I constraints.Integer, F any] interface {
	One(string) string
	Two(I) string
	Three(I) F
	Four(I) Foo[I, F]
	Five(I) Baz[F]
	Six(I) *Baz[F]
	Seven(I) other.One[I]
	Eight(F) other.Two[I, F]
	Nine(Iface[I])
	Ten(*I)
	Eleven() map[string]I
	Twelve(ctx context.Context) <-chan []I
	Thirteen(...I) *F
}

type EmbeddingIface[T constraints.Integer, R constraints.Float] interface {
	io.Reader
	Generater[R]
	Earth[Generater[T]]
	other.Either[R, StructType, other.Five, Generater[T]]
	ExternalConstraint[T, R]
}

type Generater[T any] interface {
	Generate() T
}

type Group[T Generater[T]] interface {
	Join(ctx context.Context) ([]T, error)
}
