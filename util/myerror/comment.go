package myerror

import (
	"fmt"
	"net/http"

	"git.teqnological.asia/teq-go/teq-pkg/teqerror"
)

func ErrCommentGet(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "80000",
		Message:   "Failed to get comment.",
		IsSentry:  true,
	}
}

func ErrCommentCreate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "80001",
		Message:   "Failed to create comment.",
		IsSentry:  true,
	}
}

func ErrCommentUpdate(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "80002",
		Message:   "Failed to update comment.",
		IsSentry:  true,
	}
}

func ErrCommentDelete(err error) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       err,
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: "80003",
		Message:   "Failed to delete comment.",
		IsSentry:  true,
	}
}

func ErrCommentNotFound() teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusNotFound,
		ErrorCode: "80004",
		Message:   "Not found.",
		IsSentry:  false,
	}
}

func ErrCommentInvalidParam(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusBadRequest,
		ErrorCode: "80005",
		Message:   fmt.Sprintf("Invalid paramemter: `%s`.", param),
		IsSentry:  false,
	}
}

func ErrCommentForbidden(param string) teqerror.TeqError {
	return teqerror.TeqError{
		Raw:       nil,
		HTTPCode:  http.StatusForbidden,
		ErrorCode: "80006",
		Message:   fmt.Sprintf("Denied to access the comment: `%s`.", param),
		IsSentry:  false,
	}
}
