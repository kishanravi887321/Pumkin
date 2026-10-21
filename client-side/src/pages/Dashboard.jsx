import { useState } from 'react'
import { Link } from 'react-router-dom'
import { useAuth } from '../context/AuthContext.jsx'

const NAV_ITEMS = [
  ['overview', 'Overview', '⌂'],
  ['content', 'Content studio', '▤'],
  ['calendar', 'Content calendar', '◷'],
  ['analytics', 'Analytics', '⌁'],
  ['settings', 'Workspace settings', '⚙'],
]

const METRICS = [
  ['12,480', 'Total views', '+18.2%', 'metric-card--blue'],
  ['248', 'Published posts', '+12.4%', 'metric-card--purple'],
  ['86.4%', 'Engagement rate', '+4.8%', 'metric-card--green'],
  ['34', 'Scheduled drafts', '+6.1%', 'metric-card--orange'],
  ['08', 'Team members', 'Active now', 'metric-card--slate'],
]

const CONTENT = [
  ['Arctic field notes: the 2026 survey', 'Blog post', 'Published', '12 min ago', 'blue'],
  ['A visual guide to polar cloud formations', 'Social campaign', 'In review', '43 min ago', 'purple'],
  ['March research newsletter', 'Newsletter', 'Scheduled', '2 hrs ago', 'green'],
  ['Behind the scenes: ice core lab', 'Video brief', 'Draft', 'Yesterday', 'orange'],
  ['Antarctica expedition report', 'Long-form article', 'Published', 'Yesterday', 'cyan'],
]

const ACTIVITY = [
  ['Maya Chen', 'approved “March research newsletter”', '2 minutes ago', 'MC'],
  ['You', 'created a new content brief', '18 minutes ago', 'YK'],
  ['Noah Williams', 'scheduled “Polar science weekly”', '1 hour ago', 'NW'],
  ['Ava Patel', 'commented on “Ice core lab”', '3 hours ago', 'AP'],
  ['System', 'completed your weekly analytics report', 'Yesterday', 'SY'],
]

export default function Dashboard() {
  const { user } = useAuth()
  const [activeSection, setActiveSection] = useState('overview')
  const activeItem = NAV_ITEMS.find(([id]) => id === activeSection)

  return (
    <div className="dashboard-shell">
      <aside className="dashboard-sidebar">
        <Link to="/" className="dashboard-brand">
          <span className="dashboard-brand__mark">C</span>
          <span>ContentPilot<small>WORKSPACE</small></span>
        </Link>

        <div className="dashboard-workspace">
          <span className="dashboard-workspace__label">Workspace</span>
          <button type="button" className="workspace-switcher">
            <span className="workspace-switcher__icon">CP</span>
            ContentPilot team
            <span className="workspace-switcher__chevron">⌄</span>
          </button>
        </div>

        <nav className="dashboard-nav" aria-label="Dashboard navigation">
          <span className="dashboard-nav__label">Main menu</span>
          {NAV_ITEMS.map(([id, label, icon]) => (
            <button
              type="button"
              key={id}
              className={`dashboard-nav__item${activeSection === id ? ' is-active' : ''}`}
              onClick={() => setActiveSection(id)}
            >
              <span className="dashboard-nav__icon" aria-hidden="true">{icon}</span>
              {label}
              {id === 'content' && <span className="dashboard-nav__count">5</span>}
            </button>
          ))}
          <span className="dashboard-nav__label dashboard-nav__label--secondary">Library</span>
          <Link className="dashboard-nav__item" to="/repository"><span className="dashboard-nav__icon">▥</span>Repository</Link>
          <Link className="dashboard-nav__item" to="/expeditions"><span className="dashboard-nav__icon">◎</span>Expeditions</Link>
        </nav>

        <div className="dashboard-sidebar__bottom">
          <div className="dashboard-storage"><div><span>Storage</span><strong>68%</strong></div><span className="dashboard-storage__track"><i /></span><small>6.8 GB of 10 GB used</small></div>
          <button type="button" className="dashboard-profile"><span className="avatar avatar--small">{(user?.name || 'YK').slice(0, 2).toUpperCase()}</span><span><b>{user?.name || 'Your account'}</b><small>Administrator</small></span><span className="dashboard-profile__dots">•••</span></button>
        </div>
      </aside>

      <main className="dashboard-main">
        <header className="dashboard-topbar">
          <div className="dashboard-breadcrumb"><span>Workspace</span><b>/</b><strong>{activeItem?.[1]}</strong></div>
          <div className="dashboard-topbar__actions"><button type="button" className="dashboard-icon-button" aria-label="Search">⌕</button><button type="button" className="dashboard-icon-button dashboard-notification" aria-label="Notifications">♢<i /></button><span className="dashboard-topbar__divider" /><span className="avatar">{(user?.name || 'YK').slice(0, 2).toUpperCase()}</span></div>
        </header>

        <section className="dashboard-content">
          <div className="dashboard-welcome">
            <div><span className="dashboard-kicker">Wednesday, October 21, 2026</span><h1>Good morning, {user?.name?.split(' ')[0] || 'there'} <span aria-hidden="true">✦</span></h1><p>Here is what is happening across your workspace today.</p></div>
            <button type="button" className="dashboard-primary-button">＋ Create content</button>
          </div>

          <div className="dashboard-metrics">
            {METRICS.map(([value, label, change, tone]) => <article className={`metric-card ${tone}`} key={label}><div className="metric-card__top"><span>{label}</span><span className="metric-card__icon">↗</span></div><strong>{value}</strong><small><b>↑ {change}</b> <span>vs last month</span></small></article>)}
          </div>

          <div className="dashboard-grid">
            <section className="dashboard-panel dashboard-panel--content">
              <div className="dashboard-panel__header"><div><h2>Recent content</h2><p>Your latest work across all channels.</p></div><button type="button" className="dashboard-link-button">View all <span>→</span></button></div>
              <div className="content-table">
                <div className="content-table__head"><span>Title</span><span>Type</span><span>Status</span><span>Updated</span></div>
                {CONTENT.map(([title, type, status, updated, tone]) => <div className="content-table__row" key={title}><span className="content-title"><i className={`content-dot content-dot--${tone}`} /><b>{title}</b></span><span className="content-muted">{type}</span><span><em className={`status-pill status-pill--${status.toLowerCase().replace(' ', '-')}`}>{status}</em></span><span className="content-muted">{updated}</span></div>)}
              </div>
            </section>

            <section className="dashboard-panel dashboard-panel--activity">
              <div className="dashboard-panel__header"><div><h2>Activity</h2><p>Recent workspace updates.</p></div><button type="button" className="dashboard-link-button" aria-label="More activity">•••</button></div>
              <div className="activity-list">{ACTIVITY.map(([name, action, time, initials], index) => <div className="activity-item" key={`${name}-${time}`}><span className={`avatar avatar--activity avatar--activity-${index}`}>{initials}</span><p><b>{name}</b> {action}<small>{time}</small></p></div>)}</div>
            </section>
          </div>

          <section className="dashboard-panel dashboard-quick-actions"><div className="dashboard-panel__header"><div><h2>Quick actions</h2><p>Jump back into your most-used tools.</p></div></div><div className="quick-actions__grid">{['Write an article', 'Plan a campaign', 'Review analytics', 'Invite a teammate', 'Browse repository'].map((label, index) => <button type="button" className="quick-action" key={label}><span className={`quick-action__icon quick-action__icon--${index}`}>{['✎', '◇', '⌁', '+', '▥'][index]}</span><span><b>{label}</b><small>{['Start with a blank canvas', 'Organize your next launch', 'See what is performing', 'Grow your workspace', 'Explore your library'][index]}</small></span><span className="quick-action__arrow">→</span></button>)}</div></section>
        </section>
      </main>
    </div>
  )
}
