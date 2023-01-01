package transport

// PullTokenData is a middleware that pulls token data to the context.
// func PullTokenData(next endpoint.Endpoint) endpoint.Endpoint {
// 	return func(ctx context.Context, request any) (any, error) {
// 		claims := ctx.Value(kitjwt.JWTClaimsContextKey).(jwt.MapClaims)

// 		role, ok := claims["role"].(string)
// 		if !ok {
// 			return nil, derror.ErrUnauthorized
// 		}

// 		ctx = context.WithValue(ctx, protocol.RoleContextKey, role)

// 		id, ok := claims["id"].(float64)
// 		if !ok {
// 			return nil, derror.ErrUnauthorized
// 		}

// 		switch role {
// 		case "admin":
// 			ctx = context.WithValue(ctx, protocol.AIDContextKey, protocol.AdminID(id))
// 		case "doctor":
// 			ctx = context.WithValue(ctx, protocol.DIDContextKey, protocol.DoctorID(id))
// 		case "patient":
// 			ctx = context.WithValue(ctx, protocol.PIDContextKey, protocol.PatientID(id))
// 		default:
// 			return nil, derror.ErrUnauthorized
// 		}

// 		return next(ctx, request)
// 	}
// }
