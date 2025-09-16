import React, { useState, useEffect } from 'react';
import type { Book } from './types';


const BookList = () => {
    const [books, setBooks] = useState<Book[]>([]);
    const [isLoading, setIsLoading] = useState(true);

    useEffect(() => {
        const fetchBooks = async () => {
            try {
                const response = await fetch('http://localhost:8080/books');
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }

                console.log("Response status:", response.status);

                const data = await response.json();
                setBooks(data);
            } catch (error) {
                console.error("Failed to fetch books:", error);
            } finally {
                setIsLoading(false);
            }
        };

        fetchBooks();
    }, []); // The empty dependency array is crucial

    // render logic
    if (isLoading) {
        return <div>Loading books...</div>;
    }

    if (books.length === 0) {
        return <div>No books found.</div>;
    }

    return (
        <div>
            <h2>Book Catalogue</h2>
            <ul>
                {books.map(book => (
                    <li key={book.id}>
                    <strong>{book.title}</strong> by {book.author} ({book.year})
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default BookList;