package news

import (
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *AddFeedRequest) Validate() error {
	errors := validation.ValidateStruct(r,
		validation.Field(&r.Provider, validation.Required),
		validation.Field(&r.Category, validation.Required),
		validation.Field(&r.Url, validation.Required),
		validation.Field(&r.Url, is.URL),
	)

	if errors != nil {
		return errors.(validation.Errors)
	}

	return nil
}

func (r *DeleteFeedRequest) Validate() error {
	errors := validation.ValidateStruct(r,
		validation.Field(&r.FeedId, validation.Required),
	)

	if errors != nil {
		return errors.(validation.Errors)
	}

	return nil
}

func (r *ListArticlesRequest) Validate() error {
	errors := validation.ValidateStruct(r,
		validation.Field(&r.FeedProviders, validation.Length(1, 0)),
		validation.Field(&r.PageIndex, validation.Min(0)),
		validation.Field(&r.PageSize, validation.Min(0)),
	)

	if errors != nil {
		return errors.(validation.Errors)
	}

	return nil
}

func ValidationErrToPBErrors(err error) []*ValidationError {
	var result []*ValidationError

	errors := err.(validation.Errors)
	for field, err := range errors {
		ve := &ValidationError{
			Field:   field,
			Message: err.Error(),
		}

		result = append(result, ve)
	}

	return result
}

func NewValidationError(errors []*ValidationError) error {
	status := status.New(codes.InvalidArgument, "Some of the fields in your request are not valid")
	status, _ = status.WithDetails(&ValidationErrors{Errors: errors})
	result := status.Err()
	return result
}
