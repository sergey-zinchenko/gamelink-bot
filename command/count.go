package command

import (
	"context"
	"gamelinkBot/prot"
	"gamelinkBot/service"
)

type (
	//CountFabic - struct for Count fabric
	CountFabric struct{}
	//CountCommand - struct for count command
	CountCommand struct {
		params []*prot.OneCriteriaStruct
		res    Responder
	}
)

const (
	//commandCount - const for command
	commandCount = "count"
)

//init - func for register fabric in parser
func init() {
	SharedParser().RegisterFabric(CountFabric{})
}

//RequireAdmin - func for checking if admin permissions required
func (c CountFabric) RequireAdmin() bool {
	return false
}

//Require - return array of needed permissions
func (c CountFabric) Require() []string {
	return []string{commandCount}
}

//TryParse - func for parsing request
func (c CountFabric) TryParse(req RequesterResponder) (Command, error) {
	var (
		command CountCommand
		err     error
	)
	if command.params, err = service.CompareParseCommand(req.Request(), "/"+commandCount); err != nil {
		return nil, err
	}
	command.res = req
	return command, nil
}

//Execute - execute command
func (cc CountCommand) Execute(ctx context.Context) {
	r, err := SharedClient().Count(ctx, &prot.MultiCriteriaRequest{Params: cc.params})
	if err != nil {
		cc.res.Respond(err.Error())
		return
	}
	cc.res.Respond(r.String())
}
