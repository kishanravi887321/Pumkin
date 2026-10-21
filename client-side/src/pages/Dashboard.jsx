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

const CALENDAR_ITEMS = [
  ['Mon 26', 'Polar science weekly', 'Newsletter', '09:00', 'scheduled'],
  ['Tue 27', 'The hidden life under the ice', 'Blog post', '11:30', 'scheduled'],
  ['Wed 28', 'Field notes: Weddell Sea', 'Social campaign', '14:00', 'draft'],
  ['Thu 29', 'Researcher interview — Dr. Lin', 'Video brief', '10:00', 'review'],
  ['Fri 30', 'October archive digest', 'Newsletter', '16:30', 'scheduled'],
]

const ANALYTICS = [
  ['Mon', 42, '1,820'],
  ['Tue', 68, '2,940'],
  ['Wed', 54, '2,360'],
  ['Thu', 82, '3,180'],
  ['Fri', 74, '2,790'],
  ['Sat', 91, '3,640'],
  ['Sun', 63, '2,510'],
]

function SectionHeading({ eyebrow, title, description, action }) {
  return <div className="dashboard-section-heading"><div><span className="dashboard-kicker">{eyebrow}</span><h1>{title}</h1><p>{description}</p></div>{action && <button type="button" className="dashboard-primary-button">{action}</button>}</div>
}

function Overview({ user, setActiveSection }) {
  return <>
    <div className="dashboard-welcome"><div><span className="dashboard-kicker">Wednesday, October 21, 2026</span><h1>Good morning, {user?.name?.split(' ')[0] || 'there'} <span aria-hidden="true">✦</span></h1><p>Here is what is happening across your workspace today.</p></div><button type="button" className="dashboard-primary-button" onClick={() => setActiveSection('content')}>＋ Create content</button></div>
    <div className="dashboard-metrics">{METRICS.map(([value, label, change, tone]) => <article className={`metric-card ${tone}`} key={label}><div className="metric-card__top"><span>{label}</span><span className="metric-card__icon">↗</span></div><strong>{value}</strong><small><b>↑ {change}</b> <span>vs last month</span></small></article>)}</div>
    <div className="dashboard-grid"><section className="dashboard-panel dashboard-panel--content"><PanelHeader title="Recent content" description="Your latest work across all channels." action="View all" /><ContentTable /></section><section className="dashboard-panel dashboard-panel--activity"><PanelHeader title="Activity" description="Recent workspace updates." action="•••" /><ActivityList /></section></div>
    <section className="dashboard-panel dashboard-quick-actions"><PanelHeader title="Quick actions" description="Jump back into your most-used tools." /><div className="quick-actions__grid">{['Write an article', 'Plan a campaign', 'Review analytics', 'Invite a teammate', 'Browse repository'].map((label, index) => <button type="button" className="quick-action" key={label}><span className={`quick-action__icon quick-action__icon--${index}`}>{['✎', '◇', '⌁', '+', '▥'][index]}</span><span><b>{label}</b><small>{['Start with a blank canvas', 'Organize your next launch', 'See what is performing', 'Grow your workspace', 'Explore your library'][index]}</small></span><span className="quick-action__arrow">→</span></button>)}</div></section>
  </>
}

function PanelHeader({ title, description, action }) {
  return <div className="dashboard-panel__header"><div><h2>{title}</h2><p>{description}</p></div>{action && <button type="button" className="dashboard-link-button">{action} {action === 'View all' && <span>→</span>}</button>}</div>
}

function ContentTable() {
  return <div className="content-table"><div className="content-table__head"><span>Title</span><span>Type</span><span>Status</span><span>Updated</span></div>{CONTENT.map(([title, type, status, updated, tone]) => <div className="content-table__row" key={title}><span className="content-title"><i className={`content-dot content-dot--${tone}`} /><b>{title}</b></span><span className="content-muted">{type}</span><span><em className={`status-pill status-pill--${status.toLowerCase().replace(' ', '-')}`}>{status}</em></span><span className="content-muted">{updated}</span></div>)}</div>
}

function ActivityList() {
  return <div className="activity-list">{ACTIVITY.map(([name, action, time, initials], index) => <div className="activity-item" key={`${name}-${time}`}><span className={`avatar avatar--activity avatar--activity-${index}`}>{initials}</span><p><b>{name}</b> {action}<small>{time}</small></p></div>)}</div>
}

function ContentStudio() {
  return <><SectionHeading eyebrow="Content studio" title="Create something new." description="Draft, organize, and publish content for every channel." action="＋ New content" /><div className="studio-layout"><section className="dashboard-panel studio-editor"><div className="editor-toolbar"><span className="editor-dot" /><span>Untitled content brief</span><span className="editor-toolbar__right">Saved just now&nbsp; •••</span></div><div className="editor-canvas"><span className="dashboard-kicker">CONTENT BRIEF / 0006</span><h2>What will you make today?</h2><p>Start with an idea, a research note, or a story from your archive. Your workspace is ready for the next record.</p><button type="button" className="dashboard-outline-button">Start writing&nbsp; →</button></div></section><section className="dashboard-panel studio-sidebar"><PanelHeader title="Your templates" description="Reusable starting points." />{['Research article', 'Social campaign', 'Weekly newsletter', 'Video brief', 'Expedition report'].map((item, index) => <button type="button" className="template-row" key={item}><span className={`template-icon template-icon--${index}`}>{['R', 'S', 'N', 'V', 'E'][index]}</span><span><b>{item}</b><small>Used {index + 2} times</small></span><span>→</span></button>)}</section></div></>
}

function Calendar() {
  return <><SectionHeading eyebrow="Content calendar" title="Plan the next record." description="A clear view of what is coming up across your channels." action="＋ Schedule content" /><section className="dashboard-panel calendar-panel"><div className="calendar-toolbar"><button type="button" className="dashboard-outline-button">‹</button><strong>October 26 — November 1, 2026</strong><button type="button" className="dashboard-outline-button">›</button><span className="calendar-toolbar__view">Week&nbsp; ▾</span></div><div className="calendar-list">{CALENDAR_ITEMS.map(([date, title, type, time, status]) => <div className="calendar-row" key={title}><span className="calendar-date">{date}</span><span className="content-dot content-dot--cyan" /><span className="calendar-title"><b>{title}</b><small>{type}</small></span><span className={`status-pill status-pill--${status}`}>{status === 'review' ? 'In review' : status[0].toUpperCase() + status.slice(1)}</span><time>{time}</time><button type="button" className="dashboard-link-button">•••</button></div>)}</div></section></>
}

function Analytics() {
  return <><SectionHeading eyebrow="Analytics" title="See what is resonating." description="Performance across your workspace for the last 30 days." action="Export report" /><div className="analytics-metrics">{[['12,480', 'Views', '+18.2%'], ['4,892', 'Unique visitors', '+11.8%'], ['86.4%', 'Engagement rate', '+4.8%'], ['2m 48s', 'Avg. read time', '+12.1%']].map(([value, label, change]) => <div className="analytics-stat" key={label}><span>{label}</span><strong>{value}</strong><b>↑ {change}</b></div>)}</div><section className="dashboard-panel analytics-chart"><PanelHeader title="Workspace views" description="Daily views across published content." action="Last 7 days ▾" /><div className="chart"><div className="chart-y"><span>4k</span><span>3k</span><span>2k</span><span>1k</span><span>0</span></div><div className="chart-bars">{ANALYTICS.map(([day, height, value]) => <div className="chart-bar" key={day}><span style={{ height: `${height}%` }} title={`${value} views`} /><small>{day}</small></div>)}</div></div></section><section className="dashboard-panel"><PanelHeader title="Top performing content" description="The stories your audience returned to most." /><ContentTable /></section></>
}

function Settings() {
  const settings = [['Enable short video', 'Generate a vertical short-form version from the selected content.', false], ['Create chapters', 'Analyze the transcript and add timestamped chapter sections.', true], ['Generate thumbnail', 'Generate an optional AI thumbnail based on the content and title.', false], ['Add to suitable playlist', 'Choose a relevant playlist and add published content automatically.', true]]
  return <><SectionHeading eyebrow="Workspace settings" title="Settings" description="Configure how ContentPilot works for your team." /><div className="settings-layout"><aside className="settings-tabs"><strong>Settings</strong><button type="button" className="is-selected">General</button><button type="button">Account</button><button type="button">Members</button><button type="button">Notifications</button></aside><section className="settings-content"><h2>General</h2>{settings.map(([title, description, enabled]) => <div className="setting-row" key={title}><div><h3>{title}</h3><p>{description}</p></div><button type="button" className={`toggle ${enabled ? 'is-on' : ''}`} aria-label={`Toggle ${title}`}><i /></button></div>)}<div className="setting-row setting-row--select"><div><h3>LLM model</h3><p>Choose the model used for assisted writing and summaries.</p></div><button type="button" className="settings-select">Claude 3.5 Sonnet&nbsp;⌄</button></div></section></div></>
}

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
          {activeSection === 'overview' && <Overview user={user} setActiveSection={setActiveSection} />}
          {activeSection === 'content' && <ContentStudio />}
          {activeSection === 'calendar' && <Calendar />}
          {activeSection === 'analytics' && <Analytics />}
          {activeSection === 'settings' && <Settings />}
        </section>
      </main>
    </div>
  )
}
