import fastify, { FastifyInstance } from "fastify";
import 'dotenv/config'

// Instantiate the fastify server
const server: FastifyInstance = fastify({ logger: true })

// Health check route
server.route({
    method: 'GET',
    url: '/health',
    schema: {
        response: {
            200: {
                type: 'object',
                properties: {
                    message: { type: 'string' }
                }
            }
        }
    },
    handler: async (_req, _res) => {
        return { message: 'Healthy!' }
    }
})

// Start the server
try {
    server.listen({ port: Number(process.env.PORT) ?? 5000 })
} catch (err) {
    server.log.error(err)
    process.exit(1)
} 
