import React from 'react';

interface DeleteAllProps {
  onDeleteAll: () => void;
}

const DeleteAll: React.FC<DeleteAllProps> = ({ onDeleteAll }) => {
  const handleDeleteAll = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/jobs/deleteAll', {
        method: 'DELETE',
      });
      if (!response.ok) {
        throw new Error('Failed to delete all jobs');
      }
      onDeleteAll(); // Call the parent function to update the state
    } catch (error) {
      console.error('Error deleting all jobs', error);
    }
  };

  return (
    <button 
      onClick={handleDeleteAll}
      className="bg-red-50 text-red-600 px-3 py-1 rounded-lg text-sm font-medium hover:bg-red-100 transition-colors duration-200"
    >
      Delete All Jobs
    </button>
  );
};

export default DeleteAll;
