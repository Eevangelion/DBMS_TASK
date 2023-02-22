import {React} from "react";
import { useDispatch, useSelector } from 'react-redux';
import { selectSort } from "../../store/reducers/buttons";
import "./Sorter.css";



const Sorter = (props) => {
    const userPageIsActive = useSelector(state => state.pagesReducer.userPageIsActive);
    const feedIsActive = useSelector(state => state.pagesReducer.feedIsActive);
    const searchPageIsActive = useSelector(state => state.pagesReducer.searchPageIsActive);
    const subscribesIsActive = useSelector(state => state.pagesReducer.subscribesIsActive);
    const isActive = (
        userPageIsActive &&
        feedIsActive &&
        searchPageIsActive &&
        subscribesIsActive
    );


    const dispatch = useDispatch();
    const activeButton = useSelector(state => state.buttonsReducer.sort);

    return (
        <div className="sorter" style={isActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            <button
                onClick={(event) => {
                    if (isActive)
                        dispatch(selectSort('new'))
                    else 
                        event.preventDefault();
                }}
                className={activeButton === 'new' ? 
                           'button-sort active' : 'button-sort'}
            >
                Новые
            </button>

            <div style={{marginTop: "1.6vh", marginLeft: "1.5vw", marginRight: "1.5vw", fontFamily: "Arial, Helvetica, sans-serif", fontSize: "0.7vw"}}>Лучшее за:</div>
            
            <button
                onClick={(event) => {
                    if (isActive)
                        dispatch(selectSort('hour'))
                    else 
                        event.preventDefault();
                }}
                className={activeButton === 'hour' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 час
            </button>
            <button
                onClick={(event) => {
                    if (isActive)
                        dispatch(selectSort('day'))
                    else 
                        event.preventDefault();
                }}
                className={activeButton === 'day' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 день
            </button>
            <button
                onClick={(event) => {
                    if (isActive)
                        dispatch(selectSort('week'))
                    else 
                        event.preventDefault();
                }}
                className={activeButton === 'week' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 неделю
            </button>
            <button
                onClick={(event) => {
                    if (isActive)
                        dispatch(selectSort('month'))
                    else 
                        event.preventDefault();
                }}
                className={activeButton === 'month' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 месяц
            </button>
            <button
                onClick={(event) => {
                    if (isActive)
                        dispatch(selectSort('alltime'))
                    else 
                        event.preventDefault();
                }}
                className={activeButton === 'alltime' ? 
                           'button-sort active' : 'button-sort'}
            >
                Всё время
            </button>
        
        </div>
    )
}

export default Sorter;