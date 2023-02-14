import {React} from "react";
import { useDispatch, useSelector } from 'react-redux';
import { selectSort } from "../../store/reducers/buttons";
import "./JokeSorter.css";



const JokeSorter = () => {


    const dispatch = useDispatch();
    const activeButton = useSelector(state => state.buttonsReducer.sort);

    return (
        <div className="sorter">
            <button
                onClick={() => dispatch(selectSort('new'))}
                className={activeButton === 'new' ? 
                           'button-sort active' : 'button-sort'}
            >
                Новые
            </button>

            <div style={{marginTop: "1.6vh", marginLeft: "1.5vw", marginRight: "2vw", fontFamily: "Arial, Helvetica, sans-serif", fontSize: "0.8vw"}}>Лучшее за:</div>
            
            <button
                onClick={() => dispatch(selectSort('hour'))}
                className={activeButton === 'hour' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 час
            </button>
            <button
                onClick={() => dispatch(selectSort('day'))}
                className={activeButton === 'day' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 день
            </button>
            <button
                onClick={() => dispatch(selectSort('week'))}
                className={activeButton === 'week' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 неделю
            </button>
            <button
                onClick={() => dispatch(selectSort('month'))}
                className={activeButton === 'month' ? 
                           'button-sort active' : 'button-sort'}
            >
                1 месяц
            </button>
            <button
                onClick={() => dispatch(selectSort('alltime'))}
                className={activeButton === 'alltime' ? 
                           'button-sort active' : 'button-sort'}
            >
                Всё время
            </button>
        
        </div>
    )
}

export default JokeSorter;