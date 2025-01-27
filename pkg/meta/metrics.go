/*
 * JuiceFS, Copyright 2021 Juicedata, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package meta

import "github.com/prometheus/client_golang/prometheus"

var (
	txDist = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "transaction_durations_histogram_seconds",
		Help:    "Transactions latency distributions.",
		Buckets: prometheus.ExponentialBuckets(0.0001, 1.5, 30),
	})
	txRestart = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "transaction_restart",
		Help: "The number of times a transaction is restarted.",
	})
	opDist = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "meta_ops_durations_histogram_seconds",
		Help:    "Operation latency distributions.",
		Buckets: prometheus.ExponentialBuckets(0.0001, 1.5, 30),
	})
)

func InitMetrics(registerer prometheus.Registerer) {
	if registerer == nil {
		return
	}
	registerer.MustRegister(txDist)
	registerer.MustRegister(txRestart)
	registerer.MustRegister(opDist)
}
