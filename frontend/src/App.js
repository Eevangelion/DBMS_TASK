import './App.css';
import Feed from "./pages/Feed";
import UserPage from "./pages/UserPage";
import SearchPage from "./pages/SearchPage";
import Subscribes from './pages/Subscribes';
import Settings from './components/Settings/Settings';
import TagRedactor from './components/TagRedactor/TagRedactor';
import { Route, Routes, useLocation } from 'react-router-dom';
import CreateJoke from "./components/CreateJoke/CreateJoke"
import CreateReport from './components/CreateReport/CreateReport';
import ErrorPage from './pages/ErrorPage';
import Subscribe from './components/Subscribe/Subscribe';
import ReportsList from './components/ReportsList/ReportsList';
import AuthPage from './pages/Auth';
import OAuthRedirect from './components/OAuthRedirect/OAuthRedirect';
import JokeModal from './components/JokeModal/JokeModal';
import RegisterPage from './pages/Register';
import Unsubscribe from './components/Unsubscribe/Unsubscribe';
import {AuthProvider, RequireAuth, RequireAdmin} from './context/context';
import DevelopSettings from './components/DevelopSettings/DevelopSettings';
const App = () => {
  let location = useLocation();
  let state = location.state;
  return (
    <AuthProvider>
      <Routes location={state?.backgroundLocation || location}>
        <Route index element={<RequireAuth><Feed /></RequireAuth>}/>
        <Route path='feed/' element={<RequireAuth><Feed /></RequireAuth>}/>
        <Route path='subscribes/' element={<RequireAuth><Subscribes /></RequireAuth>}/>
        <Route path='user/:username' element={<RequireAuth><UserPage /></RequireAuth>}/>
        <Route path='search/:type/:query?' element={<RequireAuth><SearchPage /></RequireAuth>}/>
        <Route path='login/' element={<AuthPage />} />
        <Route path='register/' element={<RegisterPage />} />
        <Route path='oauth/' element={<OAuthRedirect /> } />
        <Route path="tagredactor/" element={<RequireAdmin><TagRedactor /></RequireAdmin>} />
        <Route path="reportslist/" element={<RequireAdmin><ReportsList /></RequireAdmin>}/>
        <Route path="*" element={<ErrorPage />}/>
      </Routes>
      {state?.backgroundLocation && (
        <Routes>
            <Route path="/create_joke" element={<RequireAuth><CreateJoke /></RequireAuth>} />
            <Route path="/settings" element={<RequireAuth><Settings /></RequireAuth>} />
            <Route path="/develop_settings" element={<RequireAdmin><DevelopSettings/></RequireAdmin>} />
            <Route path="/create_report/:jokeID" element={<RequireAuth><CreateReport /></RequireAuth>} />
            <Route path="/subscribe/:receiverID" element={<RequireAuth><Subscribe /></RequireAuth>}/>
            <Route path="/unsubscribe/:receiverID" element={<RequireAuth><Unsubscribe /></RequireAuth>} />
            <Route path="/joke/:jokeID" element={<RequireAuth><JokeModal /></RequireAuth>} />
        </Routes>
      )}
    </AuthProvider>
  )
}

export default App;
