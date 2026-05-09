import { renderHook, act } from '@testing-library/react';
import { useIsMobile } from '@/hooks/use-mobile';

// Mock matchMedia and window.innerWidth
const mockMatchMedia = (width: number) => {
  const listeners: { [key: string]: EventListener[] } = {};
  
  Object.defineProperty(window, 'matchMedia', {
    writable: true,
    value: jest.fn().mockImplementation(query => ({
      matches: width < 768,
      media: query,
      onchange: null,
      addEventListener: jest.fn((event, listener) => {
        if (!listeners[event]) {
          listeners[event] = [];
        }
        listeners[event].push(listener as EventListener);
      }),
      removeEventListener: jest.fn(),
      dispatchEvent: jest.fn(),
    })),
  });

  Object.defineProperty(window, 'innerWidth', {
    writable: true,
    value: width,
  });
  
  // Return a function to trigger listeners
  return {
    triggerMediaChange: () => {
      if (listeners['change'] && listeners['change'].length > 0) {
        listeners['change'].forEach(listener => listener(new Event('change')));
      }
    }
  };
};

describe('useIsMobile', () => {
  beforeEach(() => {
    // Clean up mockMatchMedia between tests
    jest.clearAllMocks();
  });

  it('should return true for mobile viewport', () => {
    mockMatchMedia(767);
    const { result } = renderHook(() => useIsMobile());
    expect(result.current).toBe(true);
  });

  it('should return false for desktop viewport', () => {
    mockMatchMedia(1024);
    const { result } = renderHook(() => useIsMobile());
    expect(result.current).toBe(false);
  });

  it('should update when viewport size changes', () => {
    // Start with desktop
    const { triggerMediaChange } = mockMatchMedia(1024);
    const { result } = renderHook(() => useIsMobile());
    expect(result.current).toBe(false);

    // Change to mobile
    act(() => {
      Object.defineProperty(window, 'innerWidth', {
        writable: true,
        value: 767,
      });
      triggerMediaChange();
    });

    expect(result.current).toBe(true);
  });
}); 