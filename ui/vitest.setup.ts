import '@testing-library/jest-dom/vitest';
import { vi } from 'vitest';

// Define the mock implementation for fetch
const mockFetch = vi.fn((url: string) => {
    // You can make this smarter to handle different URLs
    if (url === 'http://localhost:8080/books') {
        return Promise.resolve({
            json: () => Promise.resolve([
                { id: '1', title: 'Test Book 1', author: 'Author A', year: 2020 },
                { id: '2', title: 'Test Book 2', author: 'Author B', year: 2021 }
            ]),
            ok: true,
        });
    }
    return Promise.reject(new Error('not mocked'));
}) as unknown as typeof fetch;

// Use vi.stubGlobal to make the mock fetch available globally in all tests
vi.stubGlobal('fetch', mockFetch);