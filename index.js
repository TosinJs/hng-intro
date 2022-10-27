const express = require("express")
const dotenv = require("dotenv")
const cors = require("cors")

dotenv.config()

const app = express()
const port = process.env.PORT || 3000

app.use(cors())

app.get("/", (req, res, next) => {
    const userDetails = {
        slackUsername: "TosinJs",
        backend: true,
        age: 98,
        bio: "I Ball, All My Guys are Ballers, We All Ball Together, Its called a Ballership"
    }
    res.json(userDetails)
})

app.listen(port, () => {
    console.log(`App Listing on Port: ${port}`)
})