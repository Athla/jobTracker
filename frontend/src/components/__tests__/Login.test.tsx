import { render, screen, fireEvent, waitFor } from "@/test/utils";
import userEvent from "@testing-library/user-event";
import Login from "../Login";
import { server } from "@/test/mocks/server";

describe("Login", () => {
  beforeAll(() => server.listen());
  afterEach(() => server.resetHandlers());
  afterAll(() => server.close());

  it("submits the form with credentials", async () => {
    render(<Login />);

    await userEvent.type(screen.getByPlaceholderText(/username/i), "testuser");
    await userEvent.type(screen.getByPlaceholderText(/password/i), "testpass");

    await userEvent.click(screen.getByRole("button", { name: /sign in/i }));

    await waitFor(() => {
      // Should redirect to dashboard after successful login
      expect(window.location.pathname).toBe("/dashboard");
    });
  });

  it("shows error message on invalid credentials", async () => {
    server.use(
      rest.post("http://localhost:8080/login", (req, res, ctx) => {
        return res(ctx.status(401), ctx.json({ error: "Invalid credentials" }));
      })
    );

    render(<Login />);

    await userEvent.type(screen.getByPlaceholderText(/username/i), "wronguser");
    await userEvent.type(screen.getByPlaceholderText(/password/i), "wrongpass");

    await userEvent.click(screen.getByRole("button", { name: /sign in/i }));

    await waitFor(() => {
      expect(
        screen.getByText(/invalid username or password/i)
      ).toBeInTheDocument();
    });
  });
});
