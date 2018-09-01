const invokeHelper = apprequire('helpers/invokeHelper.js')
const queryhelper = apprequire('helpers/queryHelper.js')

class Invoke {
  constructor() {}

  invoke(req, res) {
    let body = req.body

    invokeHelper(body)
    .then((result) => {
      res.status(200).send(result)
    })
    .catch(err => {
      res.status(400).send(err)
    })

  }
}

module.exports = Invoke
