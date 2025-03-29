import { render, screen } from '@testing-library/react';
import { AppSidebar } from '@/components/app-sidebar';
import * as React from 'react';
import { SidebarProvider } from '@/components/ui/sidebar';

// Mock the next/image component
jest.mock('next/image', () => ({
  __esModule: true,
  default: ({ src, alt, width, height }: { src: string; alt: string; width: number; height: number }) => (
    <img src={src} alt={alt} width={width} height={height} />
  ),
}));

// Mock the next/navigation functions used in AppSidebar
jest.mock('next/navigation', () => ({
  usePathname: () => '/',
}));

// Mock useIsMobile hook that is used by the SidebarProvider
jest.mock('@/hooks/use-mobile', () => ({
  useIsMobile: () => false,
}));

// Create a custom wrapper with the SidebarProvider
const SidebarWrapper = ({ children }: { children: React.ReactNode }) => (
  <SidebarProvider defaultOpen={true}>
    {children}
  </SidebarProvider>
);

// Custom render function
const customRender = (ui: React.ReactElement) => {
  return render(ui, { wrapper: SidebarWrapper });
};

describe('AppSidebar', () => {
  it('renders all navigation items', () => {
    customRender(<AppSidebar />);
    
    // Check if all navigation items are rendered
    expect(screen.getByText('Home')).toBeInTheDocument();
    expect(screen.getByText('Accounts')).toBeInTheDocument();
    expect(screen.getByText('Records')).toBeInTheDocument();
    expect(screen.getByText('Assets')).toBeInTheDocument();
    expect(screen.getByText('Statistics')).toBeInTheDocument();
    expect(screen.getByText('Budgets')).toBeInTheDocument();
  });

  it('renders the logo', () => {
    customRender(<AppSidebar />);
    
    // Check if the logo is rendered
    const logo = screen.getByAltText('Vinance Logo');
    expect(logo).toBeInTheDocument();
    expect(logo).toHaveAttribute('src', '/logo-bgless.png');
  });
}); 