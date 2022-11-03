const express = require("express")
const dotenv = require("dotenv")
const cors = require("cors")

dotenv.config()

const app = express()
const port = process.env.PORT || 3000

app.use(express.json())
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

app.post("/", (req, res, next) => {
    let { operation_type, x, y } = req.body
    if (!operation_type) {
        return res.status(400).json({error: "Bad Request", message: "Invalid Body Properties"})
    }
    const matched = operation_type.match(/\d+/g);
    if (matched && matched.length > 1) {
        x = Number(matched[0])
        y = Number(matched[1])
    }
    if (!x || !y) {
        return res.status(400).json({error: "Bad Request", message: "Invalid Body Properties"})
    }
    if (operation_type.toLowerCase().includes("add")) {
        operation_type = "addition"
    } else if (operation_type.toLowerCase().includes("subtract")) {
        operation_type = "subtraction"
    } else if (operation_type.toLowerCase().includes("multiply")) {
        operation_type = "multiplication"
    } else if (operation_type.toLowerCase().includes("multiplication")) {
        operation_type = "multiplication"
    }
    let result = 0
    switch (operation_type) {
        case "addition":
            result  = x + y
            break;
        case "subtraction":
            result = x - y
            break;
        case "multiplication":
            result = x * y
            break;
        default:
            return res.status(400).json({error: "Bad Request", message:"Invalid operation_type"})
    }
    return res.status(200).json({slackUsername: "TosinJs", operation_type: operation_type, result: result})
})

app.listen(port, () => {
    console.log(`App Listing on Port: ${port}`)
})