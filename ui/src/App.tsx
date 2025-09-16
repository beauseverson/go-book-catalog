import React from 'react';
import BookList from './components/BookList';
import './App.css'; // Optional: for styling

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>Book Catalog Application</h1>
      </header>
      <main>
        {/* Render the BookList component here */}
        <BookList />
      </main>
    </div>
  );
}

export default App;