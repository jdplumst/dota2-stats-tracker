import { FastifyInstance, FastifyServerOptions } from "fastify";

// Health check route
// export default function (
//   instance: FastifyInstance,
//   _opts: FastifyServerOptions,
//   done: any,
// ) {
//   instance.route({
//     method: "GET",
//     url: "/health",
//     schema: {
//       response: {
//         200: {
//           type: "object",
//           properties: {
//             message: { type: "string" },
//           },
//         },
//       },
//     },
//     handler: async (req, _) => {
//       req.log.info("Health check");
//       return { message: "Healthy!" };
//     },
//   });
//   done();
// }
export default async function (
  instance: FastifyInstance,
  opts: FastifyServerOptions,
  done,
) {
  instance.get("/", async (req: FastifyRequest, res: FastifyReply) => {
    res.status(200).send({
      hello: "World",
    });
  });

  instance.register(
    async (instance: FastifyInstance, opts: FastifyServerOptions, done) => {
      instance.get(
        "/",
        async (
          req: FastifyRequest<CustomRouteGenericQuery>,
          res: FastifyReply,
        ) => {
          const { name = "" } = req.query;
          res.status(200).send(`Hello ${name}`);
        },
      );

      instance.get(
        "/:name",
        async (
          req: FastifyRequest<CustomRouteGenericParam>,
          res: FastifyReply,
        ) => {
          const { name = "" } = req.params;
          res.status(200).send(`Hello ${name}`);
        },
      );
      done();
    },
    {
      prefix: "/hello",
    },
  );

  done();
}
