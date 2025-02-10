import { render, screen, fireEvent } from "@/test/utils";
import { DragDropContext } from "@hello-pangea/dnd";
import Board from "../Board";
import { Job } from "@/types/job";

const mockJobs: Job[] = [
  {
    id: "1",
    name: "Software Engineer",
    company: "Tech Corp",
    source: "LinkedIn",
    description: "Position 1",
    job_type: "FULL_TIME",
    status: "APPLIED",
    version: 1,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
  {
    id: "2",
    name: "Senior Developer",
    company: "Another Corp",
    source: "Indeed",
    description: "Position 2",
    job_type: "FULL_TIME",
    status: "PHONE_SCREEN",
    version: 1,
    created_at: "2024-01-01T00:00:00Z",
    updated_at: "2024-01-01T00:00:00Z",
  },
];

const mockOnJobMove = jest.fn();

describe("Board", () => {
  it("renders all columns", () => {
    render(<Board jobs={mockJobs} onJobMove={mockOnJobMove} />);

    expect(screen.getByText("Applied")).toBeInTheDocument();
    expect(screen.getByText("In Progress")).toBeInTheDocument();
    expect(screen.getByText("Finished")).toBeInTheDocument();
  });

  it("displays jobs in correct columns", () => {
    render(<Board jobs={mockJobs} onJobMove={mockOnJobMove} />);

    // Check if jobs are in correct columns
    const appliedJob = screen.getByText("Software Engineer");
    const inProgressJob = screen.getByText("Senior Developer");

    expect(appliedJob).toBeInTheDocument();
    expect(inProgressJob).toBeInTheDocument();
  });
});
