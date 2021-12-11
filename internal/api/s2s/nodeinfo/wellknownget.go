/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

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

package nodeinfo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/superseriousbusiness/gotosocial/internal/api"
)

// NodeInfoWellKnownGETHandler swagger:operation GET /.well-known/nodeinfo nodeInfoWellKnownGet
//
// Directs callers to /nodeinfo/2.0.
//
// eg. `{"links":[{"rel":"http://nodeinfo.diaspora.software/ns/schema/2.0","href":"http://example.org/nodeinfo/2.0"}]}`
// See: https://nodeinfo.diaspora.software/protocol.html
//
// ---
// tags:
// - nodeinfo
//
// produces:
// - application/json
//
// responses:
//   '200':
//     schema:
//       "$ref": "#/definitions/wellKnownResponse"
func (m *Module) NodeInfoWellKnownGETHandler(c *gin.Context) {
	l := logrus.WithFields(logrus.Fields{
		"func": "NodeInfoWellKnownGETHandler",
	})

	if _, err := api.NegotiateAccept(c, api.JSONAcceptHeaders...); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	niRel, err := m.processor.GetNodeInfoRel(c.Request.Context(), c.Request)
	if err != nil {
		l.Debugf("error with get node info rel request: %s", err)
		c.JSON(err.Code(), err.Safe())
		return
	}

	c.JSON(http.StatusOK, niRel)
}
