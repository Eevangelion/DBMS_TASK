import './App.css';
import Feed from "./pages/Feed";
import UserPage from "./pages/UserPage";
import SearchPage from "./pages/SearchPage";
import Subscribes from './pages/Subscribes';
import Settings from './components/Settings/Settings';
import DevelopSettings from './components/DevelopSettings/DevelopSettings';
import TagRedactor from './components/TagRedactor/TagRedactor';
import { Route, Routes, useLocation } from 'react-router-dom';
import CreateJoke from "./components/CreateJoke/CreateJoke"
import { useGetUserByIDQuery } from './services/Joke';

const App = () => {

  let location = useLocation();
  let state = location.state;
  const userID = 2;
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
        <Route path='search/?:q&:t' element={<SearchPage />}/>
      </Routes>
      {state?.backgroundLocation && (
        <Routes>
            <Route path="/create_joke" element={<CreateJoke />} />
            <Route path="/settings" element={<Settings />} />
            <Route path="/developsettings" element={<TagRedactor />} />
        </Routes>
      )}
    </>
  )
}

export default App;
