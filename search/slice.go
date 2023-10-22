package search

/* __________________________________________________ */

type Chunk struct {
	offset    int64
	offsetSet bool

	limit    int64
	limitSet bool
}

func NewChunk(offset int64, limit int64) Chunk {
	return Chunk{
		offset:    offset,
		offsetSet: true,

		limit:    limit,
		limitSet: true,
	}
}

func NewChunkWithOffset(offset int64) Chunk {
	return Chunk{
		offset:    offset,
		offsetSet: true,
	}
}

func NewChunkWithoutOffset(limit int64) Chunk {
	return Chunk{
		limit:    limit,
		limitSet: true,
	}
}

func NewChunkWithLimit(limit int64) Chunk {
	return Chunk{
		limit:    limit,
		limitSet: true,
	}
}

func NewChunkWithoutLimit(offset int64) Chunk {
	return Chunk{
		offset:    offset,
		offsetSet: true,
	}
}

func (c Chunk) Offset() (int64, bool) {
	if c.offsetSet {
		return c.offset, true
	}
	return c.offset, false
}

func (c Chunk) Limit() (int64, bool) {
	if c.limitSet {
		return c.limit, true
	}
	return c.limit, false
}

/* __________________________________________________ */

type SliceConditionOp func(*SliceCondition)

func WithChunk(chunk Chunk) SliceConditionOp {
	return func(cond *SliceCondition) {
		cond.chunk = &chunk
	}
}

func WithLimit(limit int64) SliceConditionOp {
	return func(cond *SliceCondition) {
		if cond.Chunked() {
			cond.chunk.limit = limit
			cond.chunk.limitSet = true
		}
	}
}

func WithOffset(offset int64) SliceConditionOp {
	return func(cond *SliceCondition) {
		if cond.Chunked() {
			cond.chunk.offset = offset
			cond.chunk.offsetSet = true
		}
	}
}

/* __________________________________________________ */

type SliceCondition struct {
	chunk *Chunk
}

func NewSlice() *SliceCondition {
	return &SliceCondition{}
}

func (o *SliceCondition) Chunked() bool {
	return o.chunk != nil
}

func (o *SliceCondition) Chunk() Chunk {
	if o.chunk == nil {
		panic("chunk is nil")
	}
	return *o.chunk
}

/* __________________________________________________ */
