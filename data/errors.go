package data

import (
	"database/sql"
	"errors"
)

//TODO : go thtough this file once
// Definition of all the errors that the Data Layer might return
// Inspired by database/sql
var (
	// ErrNoData To signify that no data is found for the query
	ErrNoData = errors.New("no data in result set")

	// ErrNoEntity To signify that the primary object for a relation is not found
	ErrNoEntity = errors.New("no entity found with the provided ID")

	// ErrTimedOut To signify that the no response was received from the data store in time
	ErrTimedOut = errors.New("connection to the data source has timed out")

	// ErrInvalidObject is used to signify that the object supplied was not of the correct type for the strategy used
	ErrInvalidObject = errors.New("invalid object type supplied, please check the interface implementation and supplied object")

	// ErrOperationNotSupported is used to signify that attempted usage of the data layer is invalid.
	// Invalid could indicate and inappropriate usage or indicate that the feature has yet to be implemented
	ErrOperationNotSupported = errors.New("requested operation could not be completed")

	// ErrNotSupported is used to signify not implemented functions
	ErrNotSupported = errors.New("Not Implemented")

	// ErrWrongResult is used to signify that there is a wrong number of data retrieved
	ErrWrongResult = errors.New("Wrong number of data in result set")

	// ErrRecover to signify that the data is recovered despite having an error.
	ErrRecover = errors.New("data recovered")

	allExistingErrors = []error{ErrNoData, ErrNoEntity, ErrTimedOut, ErrInvalidObject, ErrOperationNotSupported, ErrNotSupported, ErrWrongResult, ErrRecover}

	// ErrConvert is an error when converting interface.
	ErrConvert = errors.New("Error converting interface")

	// ErrPrimaryOverflow is an error when primary key overflow
	ErrPrimaryOverflow = errors.New("primary key overflow")
)

type errMsgGeneral struct {
	Message string
}

// Error implements error interface.
func (err *errMsgGeneral) Error() string {
	return err.Message
}

// ErrDuplicatedKey is used to signify the unique key is duplicated
type ErrDuplicatedKey struct {
	errMsgGeneral
}

// ErrUpdateNoResult is used to signify the update execute success but result error
type ErrUpdateNoResult struct {
	errMsgGeneral
}

// common code used when a strategy does not want to implement a particular feature
func outputNotSupported() error {
	return ErrOperationNotSupported
}

// filterNonCriticalError returns back a the same error if it is critical else returns nil
func filterNonCriticalError(err error) error {
	if err == nil {
		return err
	}

	switch err {
	case ErrNoData:
		return nil
	case sql.ErrNoRows:
		return nil
	case ErrNotSupported:
		return nil
	case ErrRecover:
		return nil
	}

	switch err.(type) {
	case *ErrUpdateNoResult:
		return nil
	case *ErrDuplicatedKey:
		return nil
	}

	return err
}

// IsUnknownDBError defines checking whether an error is not existing known error
func IsUnknownDBError(err error) bool {
	for _, existingErr := range allExistingErrors {
		if err == existingErr {
			return false
		}
	}

	if _, ok := err.(*ErrUpdateNoResult); ok {
		return false
	}

	if _, ok := err.(*ErrDuplicatedKey); ok {
		return false
	}

	return true
}
