const queryhelper = apprequire('helpers/queryHelper.js')

class Query {
  constructor() {}

  query(req, res) {
    let body = req.body
    queryhelper(body)
    .then((result) => {
      res.status(200).send(result)
    })
    .catch(err => {
      res.status(400).send(err)
    })

  }
}

module.exports = Query
