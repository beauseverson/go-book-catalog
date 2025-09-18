import React from 'react';
import { render, screen, waitFor, waitForElementToBeRemoved } from '@testing-library/react';
import { describe, it, expect } from 'vitest';
import BookList from './BookList';


describe('BookList', () => {
    it('displays a loading message, then a list of books after fetching data', async () => {
        render(<BookList />);

        expect(screen.getByText(/Loading books.../i)).toBeInTheDocument();

        await waitForElementToBeRemoved(() => screen.getByText(/Loading books.../i));

        expect(screen.getByText('Test Book 2')).toBeInTheDocument();
    });
});