# Testing Guide for Vinance Web

This project uses Jest and React Testing Library for testing.

## Getting Started

### Running Tests

To run all tests:

```bash
npm test
```

To run tests in watch mode (tests will re-run when files change):

```bash
npm run test:watch
```

To run tests with coverage report:

```bash
npm run test:coverage
```

## Test Structure

The tests are organized as follows:

```
src/
├── __tests__/
│   ├── components/    # Tests for components
│   ├── hooks/         # Tests for custom hooks
│   ├── lib/           # Tests for utility functions
│   └── test-utils.tsx # Test utilities and custom render
```

## Writing Tests

### Component Tests

Use the following pattern for testing components:

```tsx
import { render, screen } from '@testing-library/react';
// or use the custom render from test-utils if you need providers
import { render, screen } from '@/__tests__/test-utils';
import { YourComponent } from '@/components/your-component';

describe('YourComponent', () => {
  it('renders correctly', () => {
    render(<YourComponent />);
    expect(screen.getByText('Expected Text')).toBeInTheDocument();
  });
  
  // Add more tests as needed
});
```

### Hook Tests

Use the following pattern for testing hooks:

```tsx
import { renderHook, act } from '@testing-library/react';
import { useYourHook } from '@/hooks/use-your-hook';

describe('useYourHook', () => {
  it('returns the expected value', () => {
    const { result } = renderHook(() => useYourHook());
    expect(result.current).toBe(expectedValue);
  });
  
  // Add more tests as needed
});
```

## Mocking Dependencies

For components or hooks that use external dependencies (like Next.js components or browser APIs), add mocks at the top of your test file:

```tsx
// Mock Next.js components
jest.mock('next/image', () => ({
  __esModule: true,
  default: (props) => <img {...props} />,
}));

// Mock hooks
jest.mock('next/navigation', () => ({
  useRouter: () => ({
    push: jest.fn(),
    // Add other router methods as needed
  }),
}));
```

## Continuous Integration

This project uses GitHub Actions for continuous integration. Tests are automatically run on:

- Every push to the main branch
- Every pull request targeting the main branch
- Manual trigger via GitHub Actions UI

The CI workflow:
1. Builds the project
2. Runs linter checks
3. Executes all tests
4. Generates and uploads coverage reports

Coverage reports are:
- Uploaded to Codecov for tracking over time
- Available as downloadable artifacts from each GitHub Actions run

### Coverage Requirements

The project enforces minimum code coverage thresholds:
- 70% for branches
- 70% for functions
- 70% for lines
- 70% for statements

Pull requests that reduce coverage below these thresholds will fail CI checks.

## Best Practices

1. Test behavior, not implementation details
2. Use semantic queries (getByRole, getByText) instead of test IDs when possible
3. Write small, focused tests
4. Keep test code DRY but readable
5. Always add tests for new features and bug fixes
6. Always update tests when modifying existing functionality

## Resources

- [Jest Documentation](https://jestjs.io/docs/getting-started)
- [React Testing Library Documentation](https://testing-library.com/docs/react-testing-library/intro/)
- [Testing Library Cheatsheet](https://testing-library.com/docs/react-testing-library/cheatsheet/) 