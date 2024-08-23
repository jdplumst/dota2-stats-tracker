import fastify, { FastifyInstance } from "fastify";
import { env } from "./env";

// Instantiate the fastify server
const server: FastifyInstance = fastify({ logger: true });

// Health check route
server.route({
  method: "GET",
  url: "/health",
  schema: {
    response: {
      200: {
        type: "object",
        properties: {
          message: { type: "string" },
        },
      },
    },
  },
  handler: async (req, _) => {
    req.log.info("Health check");
    return { message: "Healthy!" };
  },
});

// Start the server
try {
  server.listen({ port: env.PORT ?? 5000 });
} catch (err) {
  server.log.error(err);
  process.exit(1);
}
