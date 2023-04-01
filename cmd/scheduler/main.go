
import (
	"sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)

func main() {
    command := app.NewSchedulerCommand(app.WithPlugin("ScorePlugin", ScorePlugin.New))
	
    if err := command.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
}