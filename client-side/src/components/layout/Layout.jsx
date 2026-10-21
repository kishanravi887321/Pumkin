import { Outlet, useLocation } from 'react-router-dom'
import Navbar from './Navbar.jsx'
import Footer from './Footer.jsx'

/** App shell: skip link + navbar + routed page + footer. */
export default function Layout() {
  const location = useLocation()
  const isDashboard = location.pathname === '/dashboard'

  return (
    <div className={isDashboard ? 'app-shell app-shell--dashboard' : 'app-shell'}>
      <a href="#main" className="skip-link">Skip to content</a>
      <Navbar />
      <main id="main" className="main">
        <Outlet />
      </main>
      <Footer />
    </div>
  )
}
