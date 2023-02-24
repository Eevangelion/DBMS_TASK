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
import Subscribe from './components/Subscribe/Subscribe';
import { useGetUserByIDQuery } from './services/service';
import ReportsList from './components/ReportsList/ReportsList';
import AuthModal from './components/Auth/Auth';
const App = () => {

  let location = useLocation();
  let state = location.state;
  const userID = 6;
  localStorage.setItem("userID", userID);
  const {
    data: user,
    isLoading: loadingUser,
  } = useGetUserByIDQuery(userID);
  if (!loadingUser) {
    localStorage.setItem("userName", user.name);
  }
  return (
    <>
      <Routes location={state?.backgroundLocation || location}>
        <Route index element={<Feed />}/>
        <Route path='feed/' element={<Feed />}/>
        <Route path='subscribes/' element={<Subscribes />}/>
        <Route path='user/:username' element={<UserPage />}/>
        <Route path='search/:type/:query?' element={<SearchPage />}/>
        <Route path='login/' element={<AuthModal />} />
      </Routes>
      {state?.backgroundLocation && (
        <Routes>
            <Route path="/create_joke" element={<CreateJoke />} />
            <Route path="/settings" element={<Settings />} />
            <Route path="/tagredactor" element={<TagRedactor />} />
            <Route path="/create_report/:jokeID" element={<CreateReport />} />
            <Route path="/subscribe/:receiverID" element={<Subscribe />}/>
            <Route path="/reportslist" element={<ReportsList />}/>
        </Routes>
      )}
    </>
  )
}

export default App;
