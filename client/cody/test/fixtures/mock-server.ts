import express from 'express'

// create interface for the request
interface MockRequest {
    headers: {
        authorization: string
    }
    body: {
        messages: {
            text: string
        }[]
    }
}

const SERVER_PORT = 49300

export const SERVER_URL = 'http://localhost:49300'
export const VALID_TOKEN = 'abcdefgh1234'

const responses = {
    chat: 'hello from the assistant',
    fixup: '<selection><title>Goodbye Cody</title></selection>',
}

// Runs a stub Cody service for testing.
export async function run<T>(around: () => Promise<T>): Promise<T> {
    const app = express()
    app.use(express.json())

    app.post('/.api/completions/stream', (req, res) => {
        // TODO: Filter streaming response
        // TODO: Handle multiple messages
        // Ideas from Dom - see if we could put something in the test request itself where we tell it what to respond with
        // or have a method on the server to send a set response the next time it sees a trigger word in the request.
        const request = req as MockRequest
        const response = request.body.messages[2].text.includes('<selection>') ? responses.fixup : responses.chat
        res.send(`event: completion\ndata: {"completion": ${JSON.stringify(response)}}\n\nevent: done\ndata: {}\n\n`)
    })

    app.post('/.api/graphql', (req, res) => {
        if (req.headers.authorization !== `token ${VALID_TOKEN}`) {
            res.sendStatus(401)
            return
        }

        const operation = new URL(req.url, 'https://example.com').search.replace(/^\?/, '')
        switch (operation) {
            case 'CurrentUser':
                res.send(JSON.stringify({ data: { currentUser: 'u' } }))
                break
            case 'IsContextRequiredForChatQuery':
                res.send(JSON.stringify({ data: { isContextRequiredForChatQuery: false } }))
                break
            default:
                res.sendStatus(400)
                break
        }
    })

    const server = app.listen(SERVER_PORT, () => {
        console.log(`Mock server listening on port ${SERVER_PORT}`)
    })

    const result = await around()
    server.close()

    return result
}
