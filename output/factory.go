package output

type AnalysisReporterFactory struct {
	reporter *AnalysisReporter
}

func NewAnalysisReporterFactory(archiver Archiver) *AnalysisReporterFactory {
	return &AnalysisReporterFactory{
		reporter: &AnalysisReporter{
			archiver: archiver,
		},
	}
}

func (f *AnalysisReporterFactory) AutoCloseArchiver() *AnalysisReporterFactory {
	f.reporter.closeArchiver = true
	return f
}

func (f *AnalysisReporterFactory) WithDumpStorage(ds DumpStorage) *AnalysisReporterFactory {
	f.reporter.dumpStorage = ds
	return f
}

func (f *AnalysisReporterFactory) WithFilenamePrefix(prefix string) *AnalysisReporterFactory {
	f.reporter.filenamePrefix = prefix
	return f
}

func (f *AnalysisReporterFactory) Build() *AnalysisReporter {
	return f.reporter
}
