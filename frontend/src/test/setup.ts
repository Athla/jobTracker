import "@testing-library/jest-dom";
import { server } from "./mocks/server";

/* eslint-disable @typescript-eslint/no-namespace */
declare global {
  namespace jest {
    interface Matchers<R> {
      toHaveClass: (className: string) => R;
      toBeInTheDocument: () => R;
      toHaveStyle: (style: Record<string, unknown>) => R;
    }
  }
}

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());
