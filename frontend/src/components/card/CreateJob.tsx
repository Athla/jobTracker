import React, { useState } from 'react'
 

const CreateJob: React.FC = () => {
    const [name, setName] = useState('');
    const [source, setSource] = useState('');
    const [description, setDescription] = useState('');
    const [createdAt, setCreatedAt] = useState('');
    const [isSubmitting, setIsSubmitting] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    const newJob = {
        id: crypto.randomUUID,
        name: name,
        source: source,
        description: description,
        createdat: new Date(createdAt).toISOString,
    };
    try {
        const response = await fetch('http://localhost:8080/api/jobs/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(newJob),
        })

        if (!response.ok) {
            throw new Error('Failed to create job.')
        }

        const data = await response.json();
        console.log(data)
        
        setName('');
        setSource('');
        setDescription('');
        setCreatedAt('');
        

    } catch (error) {
        console.error('Error', error)
    } finally {
        setIsSubmitting(false)
    }
    }

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label>Name:</label>
                <input type='text' value={name} onChange={(e) => setName(e.target.value)} required />
            </div>
            <div>
                <label>Source:</label>
                <input type='text' value={source} onChange={(e) => setSource(e.target.value)} required />
            </div>
            <div>
                <label>Description:</label>
                <input type='text' value={description} onChange={(e) => setDescription(e.target.value)} required />
            </div>
            <div>
                <label>Create at:</label>
                <input type='date' value={createdAt} onChange={(e) => setCreatedAt(e.target.value)} required />
            </div>
            <button type='submit' disabled={isSubmitting}>
                {isSubmitting ? 'Creating...' : 'Create Job'}
            </button>
        </form>
    )
}

export default CreateJob