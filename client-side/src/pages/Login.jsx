import { useEffect, useRef, useState } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { Alert, Button, Card, Input, PasswordInput } from '../components/ui'
import useForm from '../hooks/useForm.js'
import { login, loginWithGoogle } from '../services/authService.js'
import { compactErrors, validateEmail, validateRequiredPassword } from '../utils/validation.js'
import { useAuth } from '../context/AuthContext.jsx'

function validate(values) {
  return compactErrors({
    email: validateEmail(values.email),
    password: validateRequiredPassword(values.password),
  })
}

export default function Login() {
  const { setUser } = useAuth()
  const navigate = useNavigate()
  const location = useLocation()
  const googleButtonRef = useRef(null)
  const [googleError, setGoogleError] = useState('')
  const googleClientId = import.meta.env.VITE_GOOGLE_CLIENT_ID

  useEffect(() => {
    if (!googleClientId || !googleButtonRef.current) return undefined

    const renderGoogleButton = () => {
      if (!window.google?.accounts?.id || !googleButtonRef.current) return
      window.google.accounts.id.initialize({
        client_id: googleClientId,
        callback: async ({ credential }) => {
          try {
            setGoogleError('')
            const user = await loginWithGoogle(credential)
            setUser(user)
            navigate(location.state?.from?.pathname || '/dashboard', { replace: true })
          } catch (error) {
            setGoogleError(error.message || 'Google login failed. Please try again.')
          }
        },
      })
      googleButtonRef.current.replaceChildren()
      window.google.accounts.id.renderButton(googleButtonRef.current, {
        theme: 'outline',
        size: 'large',
        width: 360,
        text: 'signin_with',
      })
    }

    if (window.google?.accounts?.id) {
      renderGoogleButton()
      return undefined
    }

    const script = document.createElement('script')
    script.src = 'https://accounts.google.com/gsi/client'
    script.async = true
    script.onload = renderGoogleButton
    script.onerror = () => setGoogleError('Unable to load Google login.')
    document.head.appendChild(script)

    return () => {
      script.onload = null
      script.onerror = null
    }
  }, [googleClientId, location.state, navigate, setUser])

  const { submitting, submitError, handleSubmit, getFieldProps } = useForm({
    initialValues: { email: '', password: '' },
    validate,
    onSubmit: async (values) => {
      const user = await login({ email: values.email.trim(), password: values.password })
      setUser(user)
      navigate(location.state?.from?.pathname || '/dashboard', { replace: true })
    },
  })

  return (
    <div className="container auth-wrap">
      <div className="auth-panel">
        <div className="auth-header">
          <h1>Log in</h1>
          <p>Access your Pumpkin account.</p>
        </div>

        <Card>
          <form className="auth-form" onSubmit={handleSubmit} noValidate>
            {submitError && (
              <Alert variant={submitError.status === 501 ? 'info' : 'danger'}>
                {submitError.message || 'Something went wrong. Please try again.'}
              </Alert>
            )}
            {googleError && <Alert variant="danger">{googleError}</Alert>}

            <Input
              label="Email"
              type="email"
              autoComplete="email"
              inputMode="email"
              placeholder="name@example.com"
              {...getFieldProps('email')}
            />
            <PasswordInput
              label="Password"
              autoComplete="current-password"
              {...getFieldProps('password')}
            />

            <Button type="submit" fullWidth loading={submitting}>
              {submitting ? 'Logging in…' : 'Log in'}
            </Button>
          </form>

          {googleClientId ? (
            <div className="google-login">
              <span className="google-login__divider">or</span>
              <div ref={googleButtonRef} />
            </div>
          ) : (
            <Alert variant="info">
              Google login is not configured. Add VITE_GOOGLE_CLIENT_ID to client-side/.env.
            </Alert>
          )}
        </Card>

        <p className="auth-switch">
          New to Pumpkin? <Link to="/register">Create an account</Link>
        </p>
      </div>
    </div>
  )
}
