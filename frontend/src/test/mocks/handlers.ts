import { rest } from "msw";

export const handlers = [
  rest.post("http://localhost:8080/login", (req, res, ctx) => {
    const { username, password } = req.body as any;

    if (username === "testuser" && password === "testpass") {
      return res(
        ctx.status(200),
        ctx.json({
          token: "fake-jwt-token",
          type: "Bearer",
        })
      );
    }

    return res(ctx.status(401), ctx.json({ error: "Invalid credentials" }));
  }),

  rest.get("http://localhost:8080/api/jobs", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json([
        {
          id: "1",
          name: "Software Engineer",
          company: "Tech Corp",
          source: "LinkedIn",
          description: "Test job",
          status: "APPLIED",
          version: 1,
          created_at: "2024-01-01T00:00:00Z",
          updated_at: "2024-01-01T00:00:00Z",
        },
      ])
    );
  }),

  rest.patch("http://localhost:8080/api/jobs/:id/status", (req, res, ctx) => {
    const { status } = req.body as any;
    return res(
      ctx.status(200),
      ctx.json({
        id: req.params.id,
        status,
        version: 2,
        // ... other job fields
      })
    );
  }),
];
