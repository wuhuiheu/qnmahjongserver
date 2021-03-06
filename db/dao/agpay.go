// Package dao contains the types for schema 'mj'.
package dao

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// AgPay represents a row from 'mj.ag_pay'.
type AgPay struct {
	IndexID       int32     `json:"index_id"`        // index_id
	AgID          int32     `json:"ag_id"`           // ag_id
	CustomerID    int32     `json:"customer_id"`     // customer_id
	DiamondCnt    int32     `json:"diamond_cnt"`     // diamond_cnt
	MoneyCnt      int32     `json:"money_cnt"`       // money_cnt
	Delflag       int32     `json:"delflag"`         // delflag
	FirstBuyAward int32     `json:"first_buy_award"` // first_buy_award
	CreateTime    time.Time `json:"create_time"`     // create_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AgPay exists in the database.
func (ap *AgPay) Exists() bool {
	return ap._exists
}

// Deleted provides information if the AgPay has been deleted from the database.
func (ap *AgPay) Deleted() bool {
	return ap._deleted
}

// Insert inserts the AgPay to the database.
func (ap *AgPay) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ap._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO mj.ag_pay (` +
		`ag_id, customer_id, diamond_cnt, money_cnt, delflag, first_buy_award, create_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ap.AgID, ap.CustomerID, ap.DiamondCnt, ap.MoneyCnt, ap.Delflag, ap.FirstBuyAward, ap.CreateTime)
	res, err := db.Exec(sqlstr, ap.AgID, ap.CustomerID, ap.DiamondCnt, ap.MoneyCnt, ap.Delflag, ap.FirstBuyAward, ap.CreateTime)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ap.IndexID = int32(id)
	ap._exists = true

	return nil
}

// Update updates the AgPay in the database.
func (ap *AgPay) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ap._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ap._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE mj.ag_pay SET ` +
		`ag_id = ?, customer_id = ?, diamond_cnt = ?, money_cnt = ?, delflag = ?, first_buy_award = ?, create_time = ?` +
		` WHERE index_id = ?`

	// run query
	XOLog(sqlstr, ap.AgID, ap.CustomerID, ap.DiamondCnt, ap.MoneyCnt, ap.Delflag, ap.FirstBuyAward, ap.CreateTime, ap.IndexID)
	_, err = db.Exec(sqlstr, ap.AgID, ap.CustomerID, ap.DiamondCnt, ap.MoneyCnt, ap.Delflag, ap.FirstBuyAward, ap.CreateTime, ap.IndexID)
	return err
}

// Save saves the AgPay to the database.
func (ap *AgPay) Save(db XODB) error {
	if ap.Exists() {
		return ap.Update(db)
	}

	return ap.Insert(db)
}

// Delete deletes the AgPay from the database.
func (ap *AgPay) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ap._exists {
		return nil
	}

	// if deleted, bail
	if ap._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM mj.ag_pay WHERE index_id = ?`

	// run query
	XOLog(sqlstr, ap.IndexID)
	_, err = db.Exec(sqlstr, ap.IndexID)
	if err != nil {
		return err
	}

	// set deleted
	ap._deleted = true

	return nil
}

// AgPayByIndexID retrieves a row from 'mj.ag_pay' as a AgPay.
//
// Generated from index 'ag_pay_index_id_pkey'.
func AgPayByIndexID(db XODB, indexID int32) (*AgPay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_id, customer_id, diamond_cnt, money_cnt, delflag, first_buy_award, create_time ` +
		`FROM mj.ag_pay ` +
		`WHERE index_id = ?`

	// run query
	XOLog(sqlstr, indexID)
	ap := AgPay{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, indexID).Scan(&ap.IndexID, &ap.AgID, &ap.CustomerID, &ap.DiamondCnt, &ap.MoneyCnt, &ap.Delflag, &ap.FirstBuyAward, &ap.CreateTime)
	if err != nil {
		return nil, err
	}

	return &ap, nil
}

// AgPaysByAgID retrieves a row from 'mj.ag_pay' as a AgPay.
//
// Generated from index 'idx_ag_id'.
func AgPaysByAgID(db XODB, agID int32) ([]*AgPay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_id, customer_id, diamond_cnt, money_cnt, delflag, first_buy_award, create_time ` +
		`FROM mj.ag_pay ` +
		`WHERE ag_id = ?`

	// run query
	XOLog(sqlstr, agID)
	q, err := db.Query(sqlstr, agID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AgPay{}
	for q.Next() {
		ap := AgPay{
			_exists: true,
		}

		// scan
		err = q.Scan(&ap.IndexID, &ap.AgID, &ap.CustomerID, &ap.DiamondCnt, &ap.MoneyCnt, &ap.Delflag, &ap.FirstBuyAward, &ap.CreateTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &ap)
	}

	return res, nil
}

// AgPaysByAgIDCreateTime retrieves a row from 'mj.ag_pay' as a AgPay.
//
// Generated from index 'idx_ag_id_ag_id'.
func AgPaysByAgIDCreateTime(db XODB, agID int32, createTime time.Time) ([]*AgPay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_id, customer_id, diamond_cnt, money_cnt, delflag, first_buy_award, create_time ` +
		`FROM mj.ag_pay ` +
		`WHERE ag_id = ? AND create_time = ?`

	// run query
	XOLog(sqlstr, agID, createTime)
	q, err := db.Query(sqlstr, agID, createTime)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AgPay{}
	for q.Next() {
		ap := AgPay{
			_exists: true,
		}

		// scan
		err = q.Scan(&ap.IndexID, &ap.AgID, &ap.CustomerID, &ap.DiamondCnt, &ap.MoneyCnt, &ap.Delflag, &ap.FirstBuyAward, &ap.CreateTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &ap)
	}

	return res, nil
}

// AgPaysByAgIDCustomerID retrieves a row from 'mj.ag_pay' as a AgPay.
//
// Generated from index 'idx_ag_id_customer_id'.
func AgPaysByAgIDCustomerID(db XODB, agID int32, customerID int32) ([]*AgPay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_id, customer_id, diamond_cnt, money_cnt, delflag, first_buy_award, create_time ` +
		`FROM mj.ag_pay ` +
		`WHERE ag_id = ? AND customer_id = ?`

	// run query
	XOLog(sqlstr, agID, customerID)
	q, err := db.Query(sqlstr, agID, customerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AgPay{}
	for q.Next() {
		ap := AgPay{
			_exists: true,
		}

		// scan
		err = q.Scan(&ap.IndexID, &ap.AgID, &ap.CustomerID, &ap.DiamondCnt, &ap.MoneyCnt, &ap.Delflag, &ap.FirstBuyAward, &ap.CreateTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &ap)
	}

	return res, nil
}

// AgPaysByAgIDDelflag retrieves a row from 'mj.ag_pay' as a AgPay.
//
// Generated from index 'idx_ag_id_delflag'.
func AgPaysByAgIDDelflag(db XODB, agID int32, delflag int32) ([]*AgPay, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_id, customer_id, diamond_cnt, money_cnt, delflag, first_buy_award, create_time ` +
		`FROM mj.ag_pay ` +
		`WHERE ag_id = ? AND delflag = ?`

	// run query
	XOLog(sqlstr, agID, delflag)
	q, err := db.Query(sqlstr, agID, delflag)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AgPay{}
	for q.Next() {
		ap := AgPay{
			_exists: true,
		}

		// scan
		err = q.Scan(&ap.IndexID, &ap.AgID, &ap.CustomerID, &ap.DiamondCnt, &ap.MoneyCnt, &ap.Delflag, &ap.FirstBuyAward, &ap.CreateTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &ap)
	}

	return res, nil
}
