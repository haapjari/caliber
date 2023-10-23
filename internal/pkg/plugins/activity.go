package plugins

// TODO: Move this to "plugins.go"
type Plugin interface {
    MeasureActivity(string) (*ActivityReport, error)
}

type ActivityPlugin struct {

}

type ActivityReport struct {
}

// NewActivityPlugin returns ActivityPlugin.
func NewActivityPlugin() *ActivityPlugin {
    return &ActivityPlugin{}
}

// MeasureActivity returns an "ActivityReport" for a certain repository.
func (a *ActivityPlugin) MeasureActivity(path string) (*ActivityReport, error) {
    // Get Path
    // Analyze Path
    // How do you measure activity from .git repository (?)

    // TODO: This could be configurable (?)

    // Measurements:

    // Low 
   
    // Calculate total amount of commits.
    // Check the date of "first" commit.
    // Calculate amount of days, since the first commits.
     
    // First Iteration, completely without definitions. 
    // If there has not been commits, during the last "half".
    // = Low
    // If there has been commits on last month.
    // If there has not been commits during last "half".
    
    // Second Iteration, with definitions. 
    // What data can be pulled from definitions.



    // There are no recents commits. (Within 3 months.)
    // There are no commits happened last half quarter.

    // TODO: GitHub Stuff

    // There recent commits within 3 months.
    // Calculate total amount of commits.
    // Threshold of commits at least, some of the commits happened during last quarter of 
    // Threshold of commits at least 

    return nil, nil
}
