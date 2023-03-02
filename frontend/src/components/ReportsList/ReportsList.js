import { useSelector } from "react-redux";
import { useEffect, useState } from "react";
import { useGetReportsQuery } from "../../services/service";
import TopPanel from "../TopPanel/TopPanel";
import Report from "../Report/Report";
import styles from "./ReportList.module.css";
import LoadingModal from "../LoadingModal/LoadingModal";


const ReportsList = () => {
    const [pageContent, setContent] = useState(<></>);
    const reportsListIsActive = useSelector(state => state.pagesReducer.reportListIsActive);
    const {
        data: response,
        isLoading: loadingReports
    } = useGetReportsQuery();

    useEffect(() => {
        if (!loadingReports) {
            const reports = response ? response.reports : [];
            if (!reports) {
                setContent(
                    <>
                        <div className={styles.txt}>Жалоб нет</div>
                    </>
                );
            } else {
                const reportPosts = reports.map((report) => 
                {
                    return <Report report={report} />
                });
                setContent(
                    <>
                        <div className={styles.reportsList}>
                            {reportPosts}
                        </div>
                    </>
                );
            }
        }
    }, [response, loadingReports]);

    if (loadingReports) {
        return <LoadingModal />;
    }
    
    return (
        <div className={styles.mainPage}>
            <TopPanel />
            <div className={styles.info} style={reportsListIsActive ? {} : {backgroundColor: "#676a6c"}}>
                <div className={styles.feed}>
                    {pageContent}
                </div>
                {/* <PageSelector pageState={true} /> */}
            </div>
        </div>
    )
}

export default ReportsList;