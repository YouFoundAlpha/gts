/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package migrations

import (
	"context"

	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
	"github.com/uptrace/bun"
)

func init() {
	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			// set in_reply_to_uri to the uri of the status pointed to by in_reply_to_id
			if _, err := tx.NewUpdate().
				Table("statuses").
				TableExpr("statuses AS secondary").
				SetColumn("in_reply_to_uri", "secondary.uri").
				Where("statuses.in_reply_to_id = secondary.id").
				Where("statuses.in_reply_to_id IS NOT null").
				Where("statuses.in_reply_to_uri IS null").
				Exec(ctx); err != nil {
				return err
			}

			// add index to in_reply_to_uri
			if _, err := tx.
				NewCreateIndex().
				Model(&gtsmodel.Status{}).
				Index("statuses_in_reply_to_uri_idx").
				Column("in_reply_to_uri").
				Exec(ctx); err != nil {
				return err
			}

			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
