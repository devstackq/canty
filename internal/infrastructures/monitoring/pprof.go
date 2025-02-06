package monitoring

import (
	"expvar"
	"log"
	"net/http"
	_ "net/http/pprof" // Импортируем pprof для мониторинга производительности
)

// StartPerformanceMonitoring запускает мониторинг производительности
func StartPerformanceMonitoring() {
	// Настройка переменных мониторинга
	expvar.NewInt("goroutines")

	// Запуск HTTP сервера для pprof и expvar
	go func() {
		log.Println("Starting performance monitoring on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Error starting HTTP server for performance monitoring: %v", err)
		}
	}()
}
