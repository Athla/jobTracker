import {
  DragDropContext,
  Droppable,
  Draggable,
  DropResult,
  DroppableProvided,
  DroppableStateSnapshot,
} from "@hello-pangea/dnd";
import { Job, JobStatus } from "@/types/job";
import JobCard from "./JobCard";
import { cn } from "@/lib/utils";

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
    id: "finished",
    title: "Finished",
    statuses: ["OFFER", "ACCEPTED", "REJECTED", "WITHDRAWN"],
  },
];

const Board: React.FC<BoardProps> = ({ jobs, onJobMove }) => {
  const getColumnJobs = (statuses: JobStatus[]) => {
    return jobs.filter((job) => statuses.includes(job.status));
  };

  const handleDragEnd = (result: DropResult) => {
    if (!result.destination) return;

    const { draggableId, destination } = result;
    const targetColumn = columns.find(
      (col) => col.id === destination.droppableId
    );

    if (!targetColumn) return;

    // Default to the first status in the column's status list
    const newStatus = targetColumn.statuses[0];
    onJobMove(draggableId, newStatus);
  };

  return (
    <DragDropContext onDragEnd={handleDragEnd}>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {columns.map(({ id, title, statuses }) => (
          <div key={id} className="bg-card rounded-lg shadow-lg p-4">
            <h2 className="text-xl font-semibold mb-4 text-card-foreground">
              {title}
              <span className="ml-2 text-sm text-muted-foreground">
                ({getColumnJobs(statuses).length})
              </span>
            </h2>
            <Droppable droppableId={id}>
              {(
                provided: DroppableProvided,
                snapshot: DroppableStateSnapshot
              ) => (
                <div
                  {...provided.droppableProps}
                  ref={provided.innerRef}
                  className={cn(
                    "min-h-[200px] transition-colors",
                    snapshot.isDraggingOver && "bg-accent/50 rounded-lg"
                  )}
                >
                  <div className="space-y-3">
                    {getColumnJobs(statuses).map((job, index) => (
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
                              "transform transition-transform",
                              snapshot.isDragging && "scale-105"
                            )}
                          >
                            <JobCard job={job} />
                          </div>
                        )}
                      </Draggable>
                    ))}
                  </div>
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
