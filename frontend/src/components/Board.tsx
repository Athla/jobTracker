import {
  DragDropContext,
  Droppable,
  Draggable,
  DropResult,
} from "@hello-pangea/dnd";
import { Job, JobStatus } from "@/types/job";
import JobCard from "./JobCard";
import { cn } from "@/lib/utils";
import { useEffect, useState } from "react";
import { toast } from "@/hooks/use-toast";

interface BoardProps {
  jobs: Job[];
  onJobMove: (jobId: string, newStatus: JobStatus) => void;
}

interface Column {
  id: string;
  title: string;
  statuses: JobStatus[];
}

const columns: Column[] = [
  {
    id: "applied",
    title: "Applied",
    statuses: ["WISHLIST", "APPLIED"],
  },
  {
    id: "in-progress",
    title: "In Progress",
    statuses: ["PHONE_SCREEN", "TECHNICAL_INTERVIEW", "ONSITE"],
  },
  {
    id: "Offers",
    title: "Finished",
    statuses: ["OFFER", "ACCEPTED"],
  },
  {
    id: "finished",
    title: "Rejected || Withdrawn",
    statuses: ["REJECTED", "WITHDRAWN"],
  },
];

const Board: React.FC<BoardProps> = ({ jobs = [], onJobMove }) => {
  // Add default empty array
  const [organized, setOrganized] = useState<Record<string, Job[]>>({});

  useEffect(() => {
    const organizedJobs = columns.reduce((acc, column) => {
      acc[column.id] = jobs.filter((job) =>
        column.statuses.includes(job.status)
      );
      return acc;
    }, {} as Record<string, Job[]>);
    setOrganized(organizedJobs);
  }, [jobs]);

  const handleDragEnd = async (result: DropResult) => {
    const { destination, source, draggableId } = result;

    if (!destination || !source) return;

    const sourceColumn = columns.find((col) => col.id === source.droppableId);
    const destColumn = columns.find(
      (col) => col.id === destination.droppableId
    );

    if (!sourceColumn || !destColumn || sourceColumn.id === destColumn.id)
      return;

    try {
      const newStatus = destColumn.statuses[0];
      onJobMove(draggableId, newStatus);

    } catch (error) {
      toast({
        title: "Error",
        description: `Failed to update job status due: ${error}`,
        variant: "destructive",
      });
    }
  };

  return (
    <DragDropContext onDragEnd={handleDragEnd}>
      <div className="grid grid-cols-1 md:grid-cols-4 gap-3 w-full">
        {columns.map((column) => (
          <div key={column.id} className="bg-card rounded-lg shadow-lg">
            <div className="p-4 border-b border-border">
              <h2 className="text-xl font-semibold">
                {column.title}
                <span className="ml-2 text-sm text-muted-foreground">
                  ({organized[column.id]?.length || 0})
                </span>
              </h2>
            </div>
            <Droppable droppableId={column.id}>
              {(provided, snapshot) => (
                <div
                  {...provided.droppableProps}
                  ref={provided.innerRef}
                  className={cn(
                    "min-h-[500px] p-4 transition-colors",
                    snapshot.isDraggingOver && "bg-accent/50"
                  )}
                >
                  {organized[column.id]?.length === 0 ? (
                    <div className="flex items-center justify-center h-full text-muted-foreground">
                      <p>No jobs in this column</p>
                    </div>
                  ) : (
                    organized[column.id]?.map((job, index) => (
                      <Draggable
                        key={job.id}
                        draggableId={job.id}
                        index={index}
                      >
                        {(provided, snapshot) => (
                          <div
                            ref={provided.innerRef}
                            {...provided.draggableProps}
                            {...provided.dragHandleProps}
                            className={cn(
                              "mb-3",
                              snapshot.isDragging && "opacity-50"
                            )}
                          >
                            <JobCard job={job} />
                          </div>
                        )}
                      </Draggable>
                    ))
                  )}
                  {provided.placeholder}
                </div>
              )}
            </Droppable>
          </div>
        ))}
      </div>
    </DragDropContext>
  );
};
export default Board;
