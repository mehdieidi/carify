package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/schollz/progressbar/v3"

	"github.com/mehdieidi/carify/data/internal/database"
	"github.com/mehdieidi/carify/data/pkg/log"
	"github.com/mehdieidi/carify/data/protocol/derror"
	"github.com/mehdieidi/carify/data/services/car"
	"github.com/mehdieidi/carify/data/services/pcar"
	"github.com/mehdieidi/carify/data/services/preprocess"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := database.Open(
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PASSWORD"),
		),
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := database.AutoMigrate(db); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(os.Getenv("LOG_NAME"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile)

	ctx := context.Background()

	carStorage := car.NewStorage(db, logger)
	carService := car.NewService(carStorage, logger)
	preprocessStorage := preprocess.NewStorage(db, logger)
	preprocessService := preprocess.NewService(preprocessStorage, carService)
	pCarStorage := pcar.NewStorage(db, logger)
	pCarService := pcar.NewService(pCarStorage, preprocessService)

	fetchFlag := flag.Bool("fetch", false, "fetch data?")
	preprocessFlag := flag.Bool("preprocess", false, "preprocess?")
	oneHotFlag := flag.Bool("onehot", false, "one hot encode?")
	flag.Parse()

	if *fetchFlag {
		cities := []string{"1", "10", "11", "12", "13", "14", "15", "16", "1683", "1684", "1686", "1687", "1688", "1689", "1690", "1691", "1692", "1693", "1694", "1695", "1696", "1697", "1698", "1699", "17", "1700", "1701", "1702", "1703", "1706", "1707", "1708", "1709", "1710", "1711", "1712", "1713", "1714", "1715", "1716", "1717", "1718", "1719", "1720", "1721", "1722", "1723", "1724", "1725", "1726", "1727", "1728", "1729", "1730", "1731", "1732", "1733", "1734", "1735", "1736", "1737", "1738", "1739", "1740", "1741", "1742", "1743", "1744", "1745", "1746", "1747", "1748", "1749", "1750", "1751", "1752", "1753", "1754", "1755", "1756", "1757", "1758", "1759", "1760", "1761", "1762", "1763", "1764", "1765", "1766", "1767", "1768", "1769", "1770", "1771", "1772", "1773", "1774", "1775", "1776", "1777", "1778", "1779", "1780", "1781", "1782", "1783", "1784", "1785", "1786", "1787", "1788", "1789", "1790", "1791", "1792", "1793", "1794", "1795", "1796", "1797", "1798", "18", "1803", "1804", "1805", "1806", "1807", "1808", "1809", "1810", "1811", "1812", "1813", "1814", "1815", "1816", "1817", "1818", "1819", "1820", "1821", "1822", "1823", "1824", "1825", "1826", "1827", "1828", "1829", "1830", "1831", "1832", "1833", "1834", "1835", "1836", "1837", "1839", "1840", "1841", "1842", "1843", "1844", "1845", "1846", "1847", "1848", "1849", "1850", "1851", "1852", "1853", "1854", "1855", "1856", "1858", "1859", "1860", "1861", "1862", "1863", "1864", "1865", "1866", "1867", "1868", "1869", "1870", "1871", "1872", "1873", "1874", "1875", "1876", "19", "2", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "3", "30", "31", "314", "316", "317", "318", "32", "33", "34", "35", "36", "37", "38", "39", "4", "5", "6", "602", "660", "662", "663", "664", "665", "671", "7", "706", "707", "708", "709", "710", "743", "744", "745", "746", "747", "748", "749", "750", "751", "752", "753", "754", "756", "759", "760", "761", "762", "763", "764", "765", "766", "767", "768", "769", "770", "771", "772", "773", "774", "775", "776", "777", "778", "779", "780", "781", "782", "783", "784", "785", "786", "787", "788", "789", "790", "791", "792", "793", "794", "795", "796", "797", "798", "799", "8", "800", "802", "803", "804", "805", "806", "807", "808", "809", "810", "811", "812", "813", "814", "815", "816", "817", "818", "822", "823", "824", "825", "826", "827", "828", "829", "830", "831", "832", "833", "834", "835", "836", "837", "838", "839", "840", "841", "842", "843", "844", "845", "846", "847", "848", "849", "850", "851", "852", "853", "856", "857", "858", "859", "860", "861", "862", "863", "864", "865", "866", "867", "868", "869", "870", "871", "872", "873", "874"}
		brandModels := []string{"Pride 141 basic"}

		carTokens, err := carService.Search(ctx, cities, brandModels)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%+v car tokens found]\n", len(carTokens))

		fmt.Printf("\nSaving tokens to a file:\n")

		pb := progressbar.Default(int64(len(carTokens)))

		tokensFile, err := os.OpenFile("tokens", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer tokensFile.Close()

		for _, t := range carTokens {
			tokensFile.WriteString(fmt.Sprintf("%+v\n", t))
			pb.Add(1)
		}

		fmt.Printf("\nRetrieving data:\n")

		pb = progressbar.Default(int64(len(carTokens)))

		for _, t := range carTokens {
			_, err := carStorage.FindByToken(ctx, t)
			if err == nil {
				pb.Add(1)
				continue
			}
			if err != nil {
				if !errors.Is(err, derror.ErrUnknownCar) {
					pb.Add(1)

					logger.Error("main", log.ServiceLayer, "main", log.Args{log.LogErrKey: err, "car_token": t})
					continue
				}
			}

			time.Sleep(1 * time.Second)

			c, err := carService.Get(ctx, t)
			if err != nil {
				pb.Add(1)

				logger.Error("main", log.ServiceLayer, "main", log.Args{log.LogErrKey: err, "car_token": t})
				continue
			}

			_, err = carStorage.Store(ctx, c)
			if err != nil {
				pb.Add(1)

				logger.Error("main", log.ServiceLayer, "main", log.Args{log.LogErrKey: err, "car_token": t})
				continue
			}

			pb.Add(1)
		}
	}
	if *preprocessFlag {
		err := preprocessService.Year(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.Color(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.UsageKM(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.BodyStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.CashCost(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.MotorStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.RearChassisStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.FrontChassisStatus(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.InsuranceDue(ctx)
		if err != nil {
			panic(err)
		}

		err = preprocessService.GearBox(ctx)
		if err != nil {
			panic(err)
		}
	}
	if *oneHotFlag {
		pCarService.OneHotEncode(ctx)
	}
}
