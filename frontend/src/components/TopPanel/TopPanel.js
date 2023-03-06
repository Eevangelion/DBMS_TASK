import { useState } from "react";
import { FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import logo from "../../styles/img/logo_test.png";
import loop from "../../styles/img/loop_test.png";
import TopPanelButtons from "./TopPanelButtons";
import "./TopPanel.css";
import { useNavigate } from "react-router-dom";

function TopPanel(props) {
    const navigate = useNavigate();
    const [searchText, setSearchText] = useState('');
    const [searchType, setSearchType] = useState('keyword');
    const handleChange = (event) => {
        setSearchText(event.target.value);
    }
    const handleChangeType = (value) => {
        setSearchType(value);
    }

    return (
    <div className="top-panel">
        <div className="main-page-redirect">
            <a className="main-page-redirect-link" href="/">
                <img className="main-page-redirect-image" src={logo} alt=":("/>
            </a>
        </div>
        
        <div className="search-panel">
            <FormControl size="small" variant="outlined" style={{maxWidth: "10vw", maxHeight: "4.5vh", marginTop: "0.25vh", fontSize: "0.8vw"}}>
                <InputLabel>Поиск по</InputLabel>
                <Select
                    onChange={(event) => {
                            handleChangeType(event.target.value)
                        }
                    }
                    defaultValue={`keyword`}
                >
                    <MenuItem value={`keyword`}>Ключевому слову</MenuItem>
                    <MenuItem value={`tag`}>Тэгу</MenuItem>
                    <MenuItem value={`people`}>Имени пользователя</MenuItem>
                </Select>
            </FormControl>
            <input 
                type="search" 
                className="search" 
                placeholder="Поиск" 
                value={searchText}
                onChange={handleChange}
                onKeyDown={(e) => {if (e.key === 'Enter') navigate(`/search/${searchType}/?query=${searchText}`)}}
            />                    
            <a  className="search-submit"
                href={`/search/${searchType}/?query=${searchText}`} 
            >
                <img className="search-loop-image" src={loop} alt=":("/>
            </a>
        </div>
        
        <TopPanelButtons isAuth={true}/>
    </div>)
}
export default TopPanel;